/**
 * Mapping types for the schema file.
 */
export class Recipe {
	lang: Language;
	signals: Signal[];
}

export class Language {
	id: string;
	displayName: string;
}

export class Signal {
	id: string;
	displayName: string;
	apps: Sample[];
}

export class Sample {
	id: string;
	displayName: string;
	steps: Step[];
	dependencies: Dependency[];
}

export class Step {
	displayName: string;
	description: string;
	order: number;
	source: string;
	required: boolean
}

export class Dependency {
	id: string;
	version: string;
}

export const NoneLanguage: Language = { id: 'none', displayName: 'Select a language' };
export const NoneSignal: Signal = { id: 'none', displayName: 'Select a signal', apps: [] };
export const NoneSample: Sample = { id: 'none', displayName: 'Select a sample', dependencies: [], steps:[] };
