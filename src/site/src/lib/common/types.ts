/**
 * Mapping types for the schema file.
 */

export class LanguageDropDown {
	id: string;
	displayName: string;
}

export class SignalDropDown {
	id: string;
	displayName: string;
}

export class Recipe {
	id: string;
	languageId: string;
	signal: string;
	displayName: string;
	description?: string;
	sourceRoot: string;
	tags?: string[];
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
	static readonly none: LanguageDropDown = { id: 'none', displayName: 'Language' };
	static readonly csharp: LanguageDropDown = { id: 'csharp', displayName: 'C#' };
	static readonly go: LanguageDropDown = { id: 'go', displayName: 'Go' };
	static readonly java: LanguageDropDown = { id: 'java', displayName: 'Java' };
	static readonly js: LanguageDropDown = { id: 'js', displayName: 'JavaScript' };

	static readonly all: LanguageDropDown[] = [this.none, this.csharp, this.go, this.java, this.js];
}

export class Signals {
	static readonly none: SignalDropDown = { id: 'none', displayName: 'Signal' };

	static readonly all: SignalDropDown[] = [
		this.none,
		{ id: 'trace', displayName: 'Trace' },
		{ id: 'metrics', displayName: 'Metrics' },
		{ id: 'logs', displayName: 'Logs' }
	];
}

export class Recipes {
	static readonly none: Recipe = {
		id: 'none',
		languageId: 'none',
		signal: 'none',
		displayName: 'Select a sample',
		dependencies: [],
		sourceRoot: '',
		steps: []
	};
}
