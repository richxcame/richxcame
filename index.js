#!/usr/bin/env node
import chalk from 'chalk';
import boxen from 'boxen';

const { log } = console;

log(boxen('richxcame', { borderStyle: 'arrow', margin: 2, padding: 1 }));

log(
	'Hi 👋, My name is ' +
		chalk.bold.underline.blue('Baygeldi') +
		'. I ❤️  111(☘️ ) and VIM'
);

log(`
About me:
  ☕️  Coffee geek
  🥷   Coding ninja
  🎧  Music enthusiast
  🧠  Math lover
`);

log(`
Contacts:
  Email:  baygeldicholukov@gmail.com
  Phone:  +993 62726535
  Github: richxcame
`);
