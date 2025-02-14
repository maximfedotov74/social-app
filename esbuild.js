import esbuild from 'esbuild';
import { readdirSync } from 'fs';
import { join } from 'path';
import { clean } from 'esbuild-plugin-clean';

const wd = process.cwd();

const srcDir = join(wd, 'web', 'dev', 'js');
const outDir = join(wd, 'web', 'public', 'js');

const entryPoints = readdirSync(srcDir)
	.filter(file => file.endsWith('.ts') || file.endsWith('.js'))
	.map(file => join(srcDir, file));

esbuild
	.build({
		entryPoints,
		outdir: outDir,
		bundle: true,
		minify: true,
		sourcemap: true,
		format: 'esm',
		target: 'es2020',
		splitting: true,
		treeShaking: true,
		plugins: [clean({ patterns: './public/js/*' })],
	})
	.catch(() => process.exit(1));
