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

export class Languages {
	static readonly none = "none";
	static readonly csharp = "csharp";
	static readonly go = "go";
	static readonly js = "javascript";
}

export class Signals {
	static readonly none = "none";
	static readonly trace = "trace";
	static readonly metric = "metric";
	static readonly log = "log";
}

export class Samples {
	static readonly none = "none";
}

export const NoneLanguage: Language = { id: 'none', displayName: 'Select a language' };
export const NoneSignal: Signal = { id: 'none', displayName: 'Select a signal', apps: [] };
export const NoneSample: Sample = { id: 'none', displayName: 'Select a sample', dependencies: [], steps:[] };
