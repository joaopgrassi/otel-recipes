const path = require("path");
const fs = require("fs");

function findDockerfileDirectories(modifiedDirs) {
    const dockerfileDirectories = [];

    console.log("test")
    console.log(modifiedDirs);

    modifiedDirs.forEach(dir => {
      let currentPath = dir;
      const basePath = 'src'; // Most generic base path

      while (currentPath !== basePath && currentPath.includes(basePath)) {
        const dockerfileCheckPath = path.join(currentPath, 'Dockerfile');
        if (fs.existsSync(dockerfileCheckPath)) {
          dockerfileDirectories.push(currentPath); // Add the directory containing the Dockerfile
          break; // Stop once a Dockerfile is found
        }
        currentPath = path.dirname(currentPath); // Move up a directory level
      }
    });

    return dockerfileDirectories;
  }

// function findDockerfileDirectories(modifiedFiles) {
//   const dockerfileDirectories = [];

//   modifiedFiles.forEach((file) => {
//     let currentPath = path.dirname(file);
//     const basePath = "src"; // Most generic base path

//     while (currentPath !== basePath && currentPath.includes(basePath)) {
//       const dockerfileCheckPath = path.join(currentPath, "Dockerfile");
//       if (fs.existsSync(dockerfileCheckPath)) {
//         dockerfileDirectories.push(currentPath);
//         break;
//       }
//       // Go up a directory level
//       currentPath = path.dirname(currentPath);
//     }
//   });

//   return dockerfileDirectories;
// }

const generateBuildMatrix = (modifiedFiles) => {
  return findDockerfileDirectories(JSON.parse(modifiedFiles));
};

module.exports = generateBuildMatrix;
