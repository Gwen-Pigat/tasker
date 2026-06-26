import fs from 'fs';
import path from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const buildDir = path.join(__dirname, 'build');

function getHtmlFiles(dir, files = []) {
  if (!fs.existsSync(dir)) return files;
  const list = fs.readdirSync(dir);
  for (const file of list) {
    const filePath = path.join(dir, file);
    const stat = fs.statSync(filePath);
    if (stat.isDirectory()) {
      getHtmlFiles(filePath, files);
    } else if (file.endsWith('.html')) {
      files.push(filePath);
    }
  }
  return files;
}

const htmlFiles = getHtmlFiles(buildDir);

for (const filePath of htmlFiles) {
  let content = fs.readFileSync(filePath, 'utf-8');
  let scriptIndex = 0;
  
  // Match script tags that don't have a src attribute
  const regex = /<script\b[^>]*>([\s\S]*?)<\/script>/gi;
  
  content = content.replace(regex, (match, scriptContent) => {
    // If it has a src attribute, keep it
    if (match.includes('src=') || !scriptContent.trim()) {
      return match;
    }
    
    scriptIndex++;
    const htmlName = path.basename(filePath, '.html');
    const scriptFileName = `${htmlName}-init-${scriptIndex}.js`;
    const scriptPath = path.join(path.dirname(filePath), scriptFileName);
    
    // Write the inline JS to an external file
    fs.writeFileSync(scriptPath, scriptContent.trim(), 'utf-8');
    
    // Replace with a reference to the external file
    return `<script src="${scriptFileName}"></script>`;
  });
  
  fs.writeFileSync(filePath, content, 'utf-8');
}

console.log(`Successfully extracted inline scripts from ${htmlFiles.length} HTML files.`);
