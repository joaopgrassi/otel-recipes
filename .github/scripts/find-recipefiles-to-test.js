const utils = require("./utils");

const generateBuildMatrix = (modifiedFiles, rootDir) => {
  return utils.findRecipesToTest(JSON.parse(modifiedFiles), rootDir);
};

module.exports = generateBuildMatrix;
