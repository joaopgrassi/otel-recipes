var glob = require("glob");
const path = require("path");
const fs = require("fs");
const findRecipes = require("./find-recipefiles-to-test");

function findOutSyncRecipeFiles(modifiedFiles, rootDir) {
  const recipesToTest = findRecipes.findRecipesToTest(modifiedFiles, rootDir);
  let result = [];
  recipesToTest.forEach((f) => {
    let valid = true;
    const recipePath = path.dirname(f);
    let recipeObj = JSON.parse(fs.readFileSync(f, "utf8"));

    switch (recipeObj.languageId) {
      case "csharp":
        valid = checkCsharpDeps(recipeObj, recipePath);
        break;
      case "js":
        valid = checkJsDeps(recipeObj, recipePath);
        break;
      case "go":
        valid = checkGoDeps(recipeObj, recipePath);
        break;
      case "java":
        valid = checkJavaDeps(recipeObj, recipePath);
        break;
      case "python":
        valid = checkPythonDeps(recipeObj, recipePath);
        break;
      default:
        break;
    }

    // Save the recipes with out of sync deps
    if (!valid) {
      result.push(recipeObj.id);
    }
  });
  return result;
}

function checkCsharpDeps(recipe, recipePath) {
  // load the csproj file
  glob("*.csproj", { cwd: recipePath, nodir: true }, function (er, files) {
    if (files.length === 0) {
      console.log("No csproj file found in the directory");
      return false;
    }

    const csproj = fs.readFileSync(path.join(recipePath, files[0]), "utf8");

    recipe.dependencies.forEach((dep) => {
      // e.g., <PackageReference Include="OpenTelemetry" Version="1.7.0" />
      const packageName = `Include="${dep.id}"`;
      const version = `Version="${dep.version}"`;
      const found = csproj.includes(packageName) && csproj.includes(version);
      if (!found) {
        return false;
      }
    });
    return true;
  });
}

function checkJsDeps(recipe, recipePath) {
  const pkgJson = fs.readFileSync(
    path.join(recipePath, "package.json"),
    "utf8"
  );

  recipe.dependencies.forEach((dep) => {
    // e.g., "@opentelemetry/api": "^1.8.0"
    const packageName = `"${dep.id}": "${dep.version}"`;
    const found = pkgJson.includes(packageName);
    if (!found) {
      return false;
    }
  });
  return true;
}

function checkGoDeps(recipe, recipePath) {
  const goMod = fs.readFileSync(path.join(recipePath, "go.mod"), "utf8");

  recipe.dependencies.forEach((dep) => {
    // e.g., go.opentelemetry.io/otel v1.25.0
    const packageName = `${dep.id} ${dep.version}`;
    const found = goMod.includes(packageName);
    if (!found) {
      return false;
    }
  });
  return true;
}

function checkJavaDeps(recipe, recipePath) {
  // load the gradle file
  glob(
    "**/build.gradle",
    { cwd: recipePath, nodir: true },
    function (er, files) {
      if (files.length === 0) {
        console.log("No build.gradle file found in the directory");
        return false;
      }

      const gradle = fs.readFileSync(path.join(recipePath, files[0]), "utf8");

      recipe.dependencies.forEach((dep) => {
        let found = false;
        if (dep.includes("bom")) {
          // e.g., implementation platform("io.opentelemetry:opentelemetry-bom:1.36.0")
          found = gradle.includes(`${dep.id}:${dep.version}`);
        } else {
          // e.g., implementation platform("io.opentelemetry:opentelemetry-api")
          found = gradle.includes(dep.id);
        }
        if (!found) {
          return false;
        }
      });
      return true;
    }
  );
}

function checkPythonDeps(recipe, recipePath) {
  const req = fs.readFileSync(
    path.join(recipePath, "requirements.txt"),
    "utf8"
  );

  recipe.dependencies.forEach((dep) => {
    // e.g., opentelemetry-api==1.24.0
    const packageName = `${dep.id}==${dep.version}`;
    const found = pkgJson.includes(packageName);
    if (!found) {
      return false;
    }
  });
  return true;
}

module.exports = findOutSyncRecipeFiles;
