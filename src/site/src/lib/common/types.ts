/**
 * Mapping types for the schema file.
 */
export class Recipe {
  lang: Language;
  signals: Signal[]
}

export class Signal {
  type: SignalType;
  apps: Sample[];
}

export class Sample {
  type: AppType;
  dependencies: Dependency[];
  source: string;
}

export class Dependency {
  id: string;
  version: string;
}

export enum AppType {
	none,
  console,
  api,
  webapp
}

export enum SignalType {
	none,
  trace,
  metric,
  log
}

export enum Language {
	none,
  csharp,
  js,
  go,
  java,
  python
}
