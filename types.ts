class Language {
    lang: Lang;
    signals: Signal[]
}

class Signal {
    type: SignalType;
    apps: App[];
}

class App {
    type: AppType;
    dependencies: string[];
    source: string;
    recipeFile: string;
}

class Dependency {
    id: string;
    version: string;
}

enum AppType {
    console,
    api,
    webapp
}

enum SignalType {
    trace,
    metric,
    log
}

enum Lang {
    csharp,
    js,
    go,
    java,
    python
}
