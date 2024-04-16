const path = require("path");
const fs = require("fs");

function findDuplicatedRecipeIds(modifiedFiles, rootDir) {
  // get all other recipes
  allRecipeFilePaths = findAllRecipeFiles(rootDir);

  changedRecipeFilePaths = [];
  modifiedFiles.forEach((file) => {
    if (!file.includes("recipefile.json")) {
      return;
    }

    // construct the absolute path of all modified/added recipe files
    fp = path.join(rootDir, file);
    changedRecipeFilePaths.push(fp);

    // in case they were changed, removed from the all list
    allRecipeFilePaths = allRecipeFilePaths.filter((f) => f !== fp);
  });

  allRecipes = parseRecipeFiles(allRecipeFilePaths);
  toCheckRecipes = parseRecipeFiles(changedRecipeFilePaths);

  // verifies if any of the modified recipe files uses an existing Id
  invalid = [];
  toCheckRecipes.forEach((toCheck) => {
    if (allRecipes.some((r) => r.id == toCheck.id)) {
      invalid.push(`Duplicated Recipe id: ${toCheck.id}, source: ${toCheck.sourceRoot}`);
    }
  });
  return invalid;
}

/**
 *
 * @param {Array<string>} recipeFiles
 * @returns {Array<any>}
 */
function parseRecipeFiles(recipeFiles) {
  obs = [];

  recipeFiles.forEach((f) => {
    var recipeFile = JSON.parse(fs.readFileSync(f, "utf8"));
    obs.push(recipeFile);
  });
  return obs;
}

/**
 * Recursively find all recipe files.
 * @param {string} dir
 * @param {Array<string>} fileList
 * @returns {Array<string>}
 */
function findAllRecipeFiles(dir, fileList = []) {
  const files = fs.readdirSync(dir);

  files.forEach((file) => {
    const filePath = path.join(dir, file);
    const fileStat = fs.statSync(filePath);

    if (fileStat.isDirectory()) {
      findAllRecipeFiles(filePath, fileList);
    } else if (file === "recipefile.json") {
      fileList.push(filePath);
    }
  });

  return fileList;
}

const generateDuplicatedRecipeIds = (modifiedFiles, rootDir) => {
  return findDuplicatedRecipeIds(JSON.parse(modifiedFiles), rootDir);
};

module.exports = generateDuplicatedRecipeIds;
