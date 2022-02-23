/**
 * Mapping types for the schema file.
 */
export class Recipe {
	languageId: string;
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

export interface SignalDropDown {
	id: string;
	displayName: string;
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
	required: boolean;
}

export class Dependency {
	id: string;
	version: string;
}

export class Languages {
	static readonly none: Language = { id: 'none', displayName: 'Select a language' };
	static readonly csharp: Language = { id: 'csharp', displayName: 'C#' };
	static readonly go: Language = { id: 'go', displayName: 'Go' };
	static readonly java: Language = { id: 'java', displayName: 'Java' };
	static readonly js: Language = { id: 'js', displayName: 'JavaScript' };

	static readonly all: Language[] = [this.csharp, this.go, this.java, this.js];
}

export class Signals {
	static readonly none: SignalDropDown = { id: 'none', displayName: 'Select a signal' };

	static readonly all: SignalDropDown[] = [
		{ id: 'trace', displayName: 'Trace' },
		{ id: 'metric', displayName: 'Metric' },
		{ id: 'log', displayName: 'Log' }
	];
}

export class Samples {
	static readonly none: Sample = {
		id: 'none',
		displayName: 'Select a sample',
		dependencies: [],
		steps: []
	};
}
