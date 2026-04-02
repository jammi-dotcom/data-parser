const fs = require('fs');
const path = require('path');
const { promisify } = require('util');

const readDir = promisify(fs.readdir);
const readFile = promisify(fs.readFile);

async function parseFile(filePath) {
  const fileBuffer = await readFile(filePath);
  const fileContent = fileBuffer.toString();
  const rows = fileContent.split('\n');
  const data = rows.map((row) => {
    const columns = row.split(',');
    return columns.map((column) => column.trim());
  });
  return data;
}

async function parseDirectory(directoryPath) {
  const files = await readDir(directoryPath);
  const data = {};
  for (const file of files) {
    const filePath = path.join(directoryPath, file);
    const fileData = await parseFile(filePath);
    data[file] = fileData;
  }
  return data;
}

module.exports = { parseFile, parseDirectory };