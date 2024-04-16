const path = require("path");
const fs = require("fs");
const findRecipes = require("./find-recipefiles-to-test");

function findOutSyncRecipeFiles(modifiedFiles, rootDir) {
  const recipesToTest = findRecipes.findRecipesToTest(modifiedFiles, rootDir);
  let result = [];
  recipesToTest.forEach((f) => {
    const recipePath = path.dirname(f);
    let recipeObj = JSON.parse(fs.readFileSync(f, "utf8"));
    switch (recipeObj.languageId) {
      case "csharp":
        result.push(checkCsharpDeps(recipeObj, recipePath));
        break;
      case "js":
        result.push(checkJsDeps(recipeObj, recipePath));
        break;
      case "go":
        result.push(checkGoDeps(recipeObj, recipePath));
        break;
      case "java":
        result.push(checkJavaDeps(recipeObj, recipePath));
        break;
      case "python":
        result.push(checkPythonDeps(recipeObj, recipePath));
        break;
      default:
        break;
    }
  });
  return result;
}

function checkCsharpDeps(recipe, recipePath) {
  // load the csproj file
  const csprojPath = findFirstFileByExtension(recipePath, ".csproj");
  if (!csprojPath) {
    console.log("No csproj file found in the directory");
    return false;
  }

  const csproj = fs.readFileSync(csprojPath, "utf8");

  return checkInvalidDeps(
    recipe,
    (dep) => `Include="${dep.id}" Version="${dep.version}"`,
    (pkgName) => csproj.includes(pkgName)
  );
}

function checkJsDeps(recipe, recipePath) {
  const pkgJson = fs.readFileSync(
    path.join(recipePath, "package.json"),
    "utf8"
  );

  return checkInvalidDeps(
    recipe,
    (dep) => `"${dep.id}": "${dep.version}"`,
    (pkgName) => pkgJson.includes(pkgName)
  );
}

function checkGoDeps(recipe, recipePath) {
  const goMod = fs.readFileSync(path.join(recipePath, "go.mod"), "utf8");

  return checkInvalidDeps(
    recipe,
    (dep) => `${dep.id} ${dep.version}`,
    (pkgName) => goMod.includes(pkgName)
  );
}

function checkJavaDeps(recipe, recipePath) {
  // load the build.gradle file
  const gradlePath = findFirstFileByExtension(recipePath, ".gradle");
  if (!gradlePath) {
    console.log("No build.gradle file found in the directory");
    return false;
  }

  const gradle = fs.readFileSync(gradlePath, "utf8");

  invalidDeps = [];
  recipe.dependencies.forEach((dep) => {
    let found = false;
    if (dep.id.includes("bom")) {
      // e.g., implementation platform("io.opentelemetry:opentelemetry-bom:1.36.0")
      found = gradle.includes(`${dep.id}:${dep.version}`);
    } else {
      // e.g., implementation platform("io.opentelemetry:opentelemetry-api")
      found = gradle.includes(dep.id);
    }
    if (!found) {
      return invalidDeps.push(dep.id);
    }
  });
  if (invalidDeps.length === 0) {
    return null;
  }

  return {
    id: recipe.id,
    deps: invalidDeps,
  };
}

function checkPythonDeps(recipe, recipePath) {
  const req = fs.readFileSync(
    path.join(recipePath, "requirements.txt"),
    "utf8"
  );

  return checkInvalidDeps(
    recipe,
    (dep) => `${dep.id}==${dep.version}`,
    (pkgName) => req.includes(pkgName)
  );
}

function checkInvalidDeps(recipe, pkgNameFunc, includePredicate) {
  invalidDeps = [];
  recipe.dependencies.forEach((dep) => {
    const packageName = pkgNameFunc(dep);
    const found = includePredicate(packageName);
    if (!found) {
      invalidDeps.push(dep.id);
    }
  });
  if (invalidDeps.length === 0) {
    return null;
  }

  return {
    id: recipe.id,
    deps: invalidDeps,
  };
}

/**
 * Recursively searches for a file with a specific extension starting from a given directory.
 *
 * @param {string} directory The directory to start the search in.
 * @param {string} extension The desired file extension (including the dot, e.g., '.txt').
 * @return {string|null} The path to the first file found with the given extension, or null if no file is found.
 */
function findFirstFileByExtension(directory, extension) {
  let entries = fs.readdirSync(directory, { withFileTypes: true });

  for (let entry of entries) {
    let fullPath = path.join(directory, entry.name);
    if (entry.isDirectory()) {
      // Recurse into subdirectory
      let result = findFirstFileByExtension(fullPath, extension);
      if (result) {
        return result;
      }
    } else {
      if (path.extname(entry.name) === extension) {
        return fullPath;
      }
    }
  }

  return null;
}

module.exports = findOutSyncRecipeFiles;
