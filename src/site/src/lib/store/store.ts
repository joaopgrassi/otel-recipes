import { Sample, Recipe, Signal, SignalType, Language, AppType } from '$lib/common/types';
import { readable, writable, derived } from 'svelte/store';
import type { Readable } from 'svelte/store';
import data from '$lib/store/data.json';

let langs = data as unknown as Recipe[];

class LangStore {
	get allLanguages(): Readable<Recipe[]> {
		return readable(null, function start(set) {
			set(langs);
			return function stop() {};
		});
	}
}

export const store = new LangStore();
export const selectedLanguage = writable(Language.none);
export const selectedSignal = writable(SignalType.none);
export const selectedSample = writable(AppType.none);

export const languages: Readable<Language[]> = derived(
	[store.allLanguages, selectedLanguage],
	([$langStore, $selectedLanguage]) => {
		if ($selectedLanguage) {
			// if the selected language changed, clear the selected signal/sample stores
			selectedSignal.set(SignalType.none);
			selectedSample.set(AppType.none);
		}

		let langs = $langStore.map((l: Recipe) => l.lang);
		langs.unshift(Language.none);
		return langs;
	}
);

export const filteredSignals: Readable<SignalType[]> = derived(
	[store.allLanguages, selectedLanguage, selectedSignal],
	([$langStore, $selectedLanguage, $selectedSignal]) => {
		if (!$selectedLanguage) {
			return [];
		}
		if ($selectedSignal) {
			// if the signal changes, clear the selected sample
			selectedSample.set(AppType.none);
		}
		return $langStore
			.find((l: Recipe) => l.lang === $selectedLanguage)
			.signals.map((s: Signal) => s.type);
	}
);

export const filteredSamples: Readable<Sample[]> = derived(
	[store.allLanguages, selectedLanguage, selectedSignal],
	([$store, $selectedLanguage, $selectedSignal]) => {
		if (!$selectedLanguage || !$selectedSignal) {
			return [];
		}
		return $store
			.find((l: Recipe) => l.lang === $selectedLanguage)
			.signals.find((s: Signal) => s.type === $selectedSignal)
			.apps.map((app: Sample) => app);
	}
);
