const path = require("path");
const fs = require("fs");

function findRecipesToTest(modifiedFiles, rootDir) {
  // If the schema is changed, run the test on all recipe files
  if (modifiedFiles.includes("otel-recipes-schema.json")) {
    return findAllRecipeFiles(rootDir);
  }
  const recipeFilePaths = [];

  modifiedFiles.forEach((file) => {
    let currentPath = path.dirname(file);
    const basePath = "src";

    while (currentPath !== basePath && currentPath.includes(basePath)) {
      const recipeFileToCheckPath = path.join(currentPath, "recipefile.json");
      const absoluteRecipeFilePath = path.join(rootDir, recipeFileToCheckPath);
      if (
        fs.existsSync(recipeFileToCheckPath) &&
        !recipeFilePaths.includes(absoluteRecipeFilePath)
      ) {
        recipeFilePaths.push(absoluteRecipeFilePath);
        break; // Stop once a recipefile is found
      }
      currentPath = path.dirname(currentPath); // Move up a directory level
    }
  });

  return recipeFilePaths;
}

/**
 * Recursively find all recipe files.
 * @param {*} dir
 * @param {*} fileList
 * @returns
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

const generateBuildMatrix = (modifiedFiles, rootDir) => {
  return findRecipesToTest(JSON.parse(modifiedFiles), rootDir);
};

module.exports = generateBuildMatrix;
