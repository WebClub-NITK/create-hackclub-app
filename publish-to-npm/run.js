#!/usr/bin/env node

const { execFileSync } = require('child_process');
const path = require('path');
const fs = require('fs');

const banner = `
███████╗████████╗ █████╗ ██████╗ ████████╗    ██╗  ██╗ █████╗  ██████╗██╗  ██╗██╗███╗   ██╗ ██████╗ 
██╔════╝╚══██╔══╝██╔══██╗██╔══██╗╚══██╔══╝    ██║  ██║██╔══██╗██╔════╝██║ ██║ ██║████╗  ██║██╔════╝ 
███████╗   ██║   ███████║██████╔╝   ██║       ███████║███████║██║     ████    ██║██╔██╗ ██║██║  ███╗
╚════██║   ██║   ██╔══██║██╔══██║   ██║       ██╔══██║██╔══██║██║     ██╔═██║ ██║██║╚██╗██║██║   ██║
███████║   ██║   ██║  ██║██║   ██║  ██║       ██║  ██║██║  ██║╚██████╗██║  ██║██║██║ ╚████║╚██████╔╝
╚══════╝   ╚═╝   ╚═╝  ╚═╝╚═╝   ╚═╝  ╚═╝       ╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝ ╚═════╝ 
`;

// Determine platform
let binaryName;
if (process.platform === 'win32') {
  binaryName = 'cli-tool.exe';
} else if (process.platform === 'darwin') {
  binaryName = 'cli-tool-macos';
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
  console.log("\x1b[1m\x1b[36m=== HackClub NITK Starterkit ===\x1b[0m");

  // Execute the binary with all arguments passed through
  execFileSync(binaryPath, process.argv.slice(2), { stdio: 'inherit' });

  // Show banner after executing the binary
  console.log("\n", banner, "\n");

} catch (error) {
  console.error("Error executing binary:", error.message);
  process.exit(error.status || 1);
}
