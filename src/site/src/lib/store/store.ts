import { Recipe, LanguageDropDown, SignalDropDown, Languages, Signals, Recipes } from '$lib/common/types';
import { readable, writable, derived } from 'svelte/store';
import type { Readable } from 'svelte/store';
import data from '$lib/store/data.json';
import { browser } from '$app/environment';

let recipes = data as unknown as Recipe[];

const NONE: string = "none";

class LangStore {
	get allRecipes(): Readable<Recipe[]> {
		return readable(null, function start(set) {
			set(recipes);
			return function stop() {};
		});
	}
}

export const store = new LangStore();
export const selectedLanguage = writable(Languages.none);
export const selectedSignal = writable(Signals.none);
export const selectedSampleId = writable(NONE);

// For drop-downs
export const allLanguages: Readable<SignalDropDown[]> = readable(Languages.all);
export const allSignals: Readable<SignalDropDown[]> = readable(Signals.all);

export const filteredSamples: Readable<Recipe[]> = derived(
	[store.allRecipes, selectedLanguage, selectedSignal],
	([$store, $selectedLanguage, $selectedSignal]) => {
		// reset the query params when the selected language/signal changes
		// it's added again when the user selects a sample
		clearQueryParams();

		let filteredSamples: Recipe[] = [];

		if ($selectedLanguage.id === Languages.none.id && $selectedSignal.id === Signals.none.id) {
			return [];
		}

		// Only language filter selected
		if ($selectedLanguage.id !== Languages.none.id && $selectedSignal.id === Signals.none.id) {
			const recipes = $store.filter((r: Recipe) => r.languageId === $selectedLanguage.id);
			return recipes;
		}

		// Only signal filter selected
		if ($selectedLanguage.id === Languages.none.id && $selectedSignal.id !== Signals.none.id) {
			const recipes = $store.filter((r: Recipe) => r.signal === $selectedSignal.id);
			return recipes;
		}

		// Both filters selected
		return $store.filter((r: Recipe) => r.languageId === $selectedLanguage.id && r.signal === $selectedSignal.id);

		// let signal = $store
		// 	.find((l: Recipe) => l.languageId === $selectedLanguage.id)
		// 	.signals.find((s: Signal) => s.id === $selectedSignal.id);

		// if (!signal) {
		// 	return [];
		// }

		// let samples = signal.samples.map((app: Recipe) => app);
		// samples.unshift(Samples.none);
		// return samples;
	}
);

export const selectedSample: Readable<Recipe> = derived(
	[store.allRecipes, selectedLanguage, selectedSignal, selectedSampleId],
	([$store, $selectedLanguage, $selectedSignal, $selectedSampleId]) => {
		// if (
		// 	$selectedLanguage.id === Languages.none.id ||
		// 	$selectedSignal.id === Signals.none.id ||
		// 	$selectedSampleId === NONE
		// ) {
		// 	return NONE;
		// }

		// const recipe = $store.find((l: Recipe) => l.languageId === $selectedLanguage.id);
		// if (!recipe) {
		// 	return NONE;
		// }

		// const signal = recipe.signals.find((s: Signal) => s.id === $selectedSignal.id);
		// if (!signal) {
		// 	return Samples.none;
		// }

		// const sample = signal.samples.find((app: Recipe) => app.id === $selectedSampleId);
		// if (!sample) {
		// 	return Samples.none;
		// }

		// replaceStateWithQuery({
		// 	language: $selectedLanguage.id,
		// 	signal: $selectedSignal.id,
		// 	sample: sample.id
		// });
		// return sample;
		return Recipes.none;
	}
);

export function setFromUrl(languageId?: string, signalId?: string, sampleId?: string) {
	if (!languageId || !signalId || !sampleId) {
		return;
	}

	const language = Languages.all.find((l: LanguageDropDown) => l.id === languageId);
	if (!language) {
		// if the selected language does not exist in the list, set to none
		// This can happen for ex if someone changes the URL. E.g. /recipes?language=madeup-language
		selectedLanguage.set(Languages.none);
		return;
	}
	selectedLanguage.set(language);

	const signal = Signals.all.find((s: SignalDropDown) => s.id === signalId);
	if (!signal) {
		// if the selected signal does not exist for the language return empty
		// This can happen for ex if someone changes the URL. E.g. recipes/csharp/madeup-signal
		selectedSignal.set(Signals.none);
		return;
	}
	selectedSignal.set(signal);
	selectedSampleId.set(sampleId);
}

function clearQueryParams(): void {
	replaceStateWithQuery({
		language: null,
		signal: null,
		sample: null
	});
}

function replaceStateWithQuery(values: Record<string, string>): void {
	if (!browser) {
		return;
	}
	const url = new URL(window.location.toString());
	for (let [k, v] of Object.entries(values)) {
		if (!!v) {
			url.searchParams.set(encodeURIComponent(k), encodeURIComponent(v));
		} else {
			url.searchParams.delete(k);
		}
	}
	history.replaceState({}, '', url);
}
