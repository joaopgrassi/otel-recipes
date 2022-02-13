import {
	Sample,
	Recipe,
	Signal,
	Language,
	NoneLanguage,
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
export const selectedSample = writable(Samples.none);

export const languages: Readable<Language[]> = derived(
	[store.allLanguages, selectedLanguage],
	([$langStore, $selectedLanguage]) => {
		let langs = $langStore.map((l: Recipe) => l.lang);
		langs.unshift(NoneLanguage);

		// if the selected language does not exist in the list, set to none
		// This can happen for ex if someone changes the URL. E.g. /recipes/madeup-lang
		if (!$langStore.some((r: Recipe) => r.lang.id === $selectedLanguage)) {
			selectedLanguage.set(Languages.none);
		}

		return langs;
	}
);

export const filteredSignals: Readable<Signal[]> = derived(
	[store.allLanguages, selectedLanguage, selectedSignal],
	([$langStore, $selectedLanguage, $selectedSignal]) => {
		if ($selectedLanguage === Languages.none) {
			return [];
		}

		// if the selected language does not exist in the list, return empty
		// This can happen for ex if someone changes the URL. E.g. /recipes/madeup-lang
		if (!$langStore.some((r: Recipe) => r.lang.id === $selectedLanguage)) {
			return [];
		}

		let signals = $langStore
			.find((r: Recipe) => r.lang.id === $selectedLanguage)
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
		if ($selectedLanguage === NoneLanguage.id || $selectedSignal === NoneSignal.id) {
			return [];
		}

		let signal = $store
			.find((l: Recipe) => l.lang.id === $selectedLanguage)
			.signals.find((s: Signal) => s.id === $selectedSignal);

		if (!signal) {
			return [];
		}

		let samples = signal.apps.map((app: Sample) => app);
		samples.unshift(NoneSample);
		return samples;
	}
);
