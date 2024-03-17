const path = require("path");
const fs = require("fs");

function findDockerfileDirectories(modifiedDirs) {
    const dockerfileDirectories = [];

    modifiedDirs.forEach(dir => {
      let currentPath = dir;
      const basePath = 'src';

      while (currentPath !== basePath && currentPath.includes(basePath)) {
        const dockerfileCheckPath = path.join(currentPath, 'Dockerfile');
        if (fs.existsSync(dockerfileCheckPath) && !dockerfileDirectories.includes(currentPath)) {
          dockerfileDirectories.push(currentPath);
          break; // Stop once a Dockerfile is found
        }
        currentPath = path.dirname(currentPath); // Move up a directory level
      }
    });

    return dockerfileDirectories;
  }

const generateBuildMatrix = (modifiedFiles) => {
  return findDockerfileDirectories(JSON.parse(modifiedFiles));
};

module.exports = generateBuildMatrix;
