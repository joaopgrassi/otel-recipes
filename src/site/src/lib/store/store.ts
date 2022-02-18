import {
	Sample,
	Recipe,
	Signal,
	Language,
	NoneSignal,
	NoneSample,
	Languages,
	Signals,
	Samples
} from '$lib/common/types';
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
export const selectedSampleId = writable(Samples.none);

export const languages: Readable<Language[]> = derived(
	[store.allLanguages], ([$langStore]) => {
		// The set of languages we have samples for
		const langIds = new Set($langStore.map((r: Recipe) => r.languageId));

		const langs = Languages.all.filter((l: Language) => langIds.has(l.id));
		langs.unshift(Languages.none);

		return langs;
	}
);

export const filteredSignals: Readable<Signal[]> = derived(
	[store.allLanguages, selectedLanguage, selectedSignal],
	([$langStore, $selectedLanguage, $selectedSignal]) => {
		if ($selectedLanguage.id === Languages.none.id) {
			return [];
		}

		// if the selected language does not exist in the list, return empty
		// This can happen for ex if someone changes the URL. E.g. /recipes/madeup-lang
		if (!$langStore.some((r: Recipe) => r.languageId === $selectedLanguage.id)) {
			return [];
		}

		let signals = $langStore
			.find((r: Recipe) => r.languageId === $selectedLanguage.id)
			.signals.map((s: Signal) => s);

		signals.unshift(NoneSignal);

		// if the selected signal does not exist for the language return empty
		// This can happen for ex if someone changes the URL. E.g. recipes/csharp/madeup-signal
		if (!signals.some((s: Signal) => s.id === $selectedSignal)) {
			return [];
		}

		return signals;
	}
);

export const filteredSamples: Readable<Sample[]> = derived(
	[store.allLanguages, selectedLanguage, selectedSignal],
	([$store, $selectedLanguage, $selectedSignal]) => {
		if ($selectedLanguage.id === Languages.none.id || $selectedSignal === NoneSignal.id) {
			return [];
		}

		let signal = $store
			.find((l: Recipe) => l.languageId === $selectedLanguage.id)
			.signals.find((s: Signal) => s.id === $selectedSignal);

		if (!signal) {
			return [];
		}

		let samples = signal.apps.map((app: Sample) => app);
		samples.unshift(NoneSample);
		return samples;
	}
);

export const selectedSample: Readable<Sample> = derived(
	[store.allLanguages, selectedLanguage, selectedSignal, selectedSampleId],
	([$store, $selectedLanguage, $selectedSignal, $selectedSampleId]) => {
		if (
			$selectedLanguage.id === Languages.none.id ||
			$selectedSignal === Signals.none ||
			$selectedSampleId == Samples.none
		) {
			return NoneSample;
		}

		const recipe = $store.find((l: Recipe) => l.languageId === $selectedLanguage.id);
		if (!recipe) {
			return NoneSample;
		}

		const signal = recipe.signals.find((s: Signal) => s.id === $selectedSignal);
		if (!signal) {
			return NoneSample;
		}

		const sample = signal.apps.find((app: Sample) => app.id === $selectedSampleId);
		if (!sample) {
			return NoneSample;
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
