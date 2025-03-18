#!/usr/bin/env node

const { execFileSync } = require('child_process');
const path = require('path');
const fs = require('fs');

// Determine platform
let binaryName;
if (process.platform === 'win32') {
  binaryName = 'cli-tool-win.exe';
} else if (process.platform === 'darwin') {
  binaryName = 'cli-tool-mac.txt';
} else if (process.platform === 'linux') {
  binaryName = 'cli-tool-linux';
} else {
  console.error('Unsupported platform');
  process.exit(1);
}

const binaryPath = path.join(__dirname, 'bin', binaryName);

// Make sure the binary is executable (not needed for Windows)
if (process.platform !== 'win32') {
  fs.chmodSync(binaryPath, '755');
}

try {
  // Execute the binary with all arguments passed through
  execFileSync(binaryPath, process.argv.slice(2), { stdio: 'inherit' });
} catch (error) {
  // Forward the exit code
  process.exit(error.status);
}