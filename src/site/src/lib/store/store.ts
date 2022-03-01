import {
	Sample,
	Recipe,
	Signal,
	Language,
	Languages,
	Signals,
	Samples
} from '$lib/common/types';
import type { SignalDropDown } from '$lib/common/types';
import { readable, writable, derived } from 'svelte/store';
import type { Readable } from 'svelte/store';
import data from '$lib/store/data.json';

let recipes = data as unknown as Recipe[];

class LangStore {
	get allLanguages(): Readable<Recipe[]> {
		return readable(null, function start(set) {
			set(recipes);
			return function stop() {};
		});
	}
}

export const store = new LangStore();
export const selectedLanguage = writable(Languages.none);
export const selectedSignal = writable(Signals.none);
export const selectedSampleId = writable(Samples.none.id);

export const languages: Readable<Language[]> = derived([store.allLanguages], ([$langStore]) => {
	// The set of languages we have samples for
	const langIds = new Set($langStore.map((r: Recipe) => r.languageId));

	const langs = Languages.all.filter((l: Language) => langIds.has(l.id));
	langs.unshift(Languages.none);

	return langs;
});

export const filteredSignals: Readable<SignalDropDown[]> = derived(
	[store.allLanguages, selectedLanguage],
	([$langStore, $selectedLanguage]) => {
		if ($selectedLanguage.id === Languages.none.id) {
			return [];
		}

		// if the selected language does not exist in the list, return empty
		// This can happen for ex if someone changes the URL. E.g. /recipes/madeup-lang
		if (!$langStore.some((r: Recipe) => r.languageId === $selectedLanguage.id)) {
			return [];
		}

		const signalsIds = new Set(
			$langStore
				.find((r: Recipe) => r.languageId === $selectedLanguage.id)
				.signals.map((s: Signal) => s.id)
		);

		const signals = Signals.all.filter((s: SignalDropDown) => signalsIds.has(s.id));
		signals.unshift(Signals.none);

		return signals;
	}
);

export const filteredSamples: Readable<Sample[]> = derived(
	[store.allLanguages, selectedLanguage, selectedSignal],
	([$store, $selectedLanguage, $selectedSignal]) => {
		if ($selectedLanguage.id === Languages.none.id || $selectedSignal.id === Signals.none.id) {
			return [];
		}

		let signal = $store
			.find((l: Recipe) => l.languageId === $selectedLanguage.id)
			.signals.find((s: Signal) => s.id === $selectedSignal.id);

		if (!signal) {
			return [];
		}

		let samples = signal.samples.map((app: Sample) => app);
		samples.unshift(Samples.none);
		return samples;
	}
);

export const selectedSample: Readable<Sample> = derived(
	[store.allLanguages, selectedLanguage, selectedSignal, selectedSampleId],
	([$store, $selectedLanguage, $selectedSignal, $selectedSampleId]) => {
		if (
			$selectedLanguage.id === Languages.none.id ||
			$selectedSignal.id === Signals.none.id ||
			$selectedSampleId == Samples.none.id
		) {
			return Samples.none;
		}

		const recipe = $store.find((l: Recipe) => l.languageId === $selectedLanguage.id);
		if (!recipe) {
			return Samples.none;
		}

		const signal = recipe.signals.find((s: Signal) => s.id === $selectedSignal.id);
		if (!signal) {
			return Samples.none;
		}

		const sample = signal.samples.find((app: Sample) => app.id === $selectedSampleId);
		if (!sample) {
			return Samples.none;
		}

		return sample;
	}
);

export function setSelectedLanguage(langId: string): void {
	const lang = Languages.all.find((l: Language) => l.id === langId);

	if (lang) {
		selectedLanguage.set(lang);
		return;
	}

	// if the selected language does not exist in the list, set to none
	// This can happen for ex if someone changes the URL. E.g. /recipes/madeup-lang
	selectedLanguage.set(Languages.none);
}

export function setSelectedSignal(signalId: string): void {
	const signal = Signals.all.find((s: SignalDropDown) => s.id === signalId);

	if (signal) {
		selectedSignal.set(signal);
		return;
	}

	// if the selected signal does not exist for the language return empty
	// This can happen for ex if someone changes the URL. E.g. recipes/csharp/madeup-signal
	selectedSignal.set(Signals.none);
}
