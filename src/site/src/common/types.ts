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
  console,
  api,
  webapp
}

export enum SignalType {
  trace,
  metric,
  log
}

export enum Language {
  csharp,
  js,
  go,
  java,
  python
}
