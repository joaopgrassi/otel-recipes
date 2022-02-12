import {
	Sample,
	Recipe,
	Signal,
	Language,
	NoneLanguage,
	NoneSignal,
	NoneSample
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
export const selectedLanguage = writable(NoneLanguage);
export const selectedSignal = writable(NoneSignal);
export const selectedSample = writable(NoneSample);

export const languages: Readable<Language[]> = derived(
	[store.allLanguages, selectedLanguage],
	([$langStore, $selectedLanguage]) => {
		if ($selectedLanguage) {
			// if the selected language changed, clear the selected signal/sample stores
			selectedSignal.set(NoneSignal);
			selectedSample.set(NoneSample);
		}

		let langs = $langStore.map((l: Recipe) => l.lang);
		langs.unshift(NoneLanguage);
		return langs;
	}
);

export const filteredSignals: Readable<Signal[]> = derived(
	[store.allLanguages, selectedLanguage, selectedSignal],
	([$langStore, $selectedLanguage, $selectedSignal]) => {
		if ($selectedLanguage.id == NoneLanguage.id) {
			return [];
		}
		if ($selectedSignal) {
			// if the signal changes, clear the selected sample
			selectedSample.set(NoneSample);
		}

		let signals = $langStore
			.find((l: Recipe) => l.lang.id === $selectedLanguage.id)
			.signals.map((s: Signal) => s);

		signals.unshift(NoneSignal);
		return signals;
	}
);

export const filteredSamples: Readable<Sample[]> = derived(
	[store.allLanguages, selectedLanguage, selectedSignal],
	([$store, $selectedLanguage, $selectedSignal]) => {
		if ($selectedLanguage.id == NoneLanguage.id || $selectedSignal.id == NoneSignal.id) {
			return [];
		}

		let samples = $store
			.find((l: Recipe) => l.lang.id === $selectedLanguage.id)
			.signals.find((s: Signal) => s.id === $selectedSignal.id)
			.apps.map((app: Sample) => app);

		samples.unshift(NoneSample);
		return samples;
	}
);
