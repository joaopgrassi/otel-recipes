import {
	Recipe,
	LanguageDropDown,
	SignalDropDown,
	Languages,
	Signals,
	Recipes
} from '$lib/common/types';
import { readable, writable, derived } from 'svelte/store';
import type { Readable } from 'svelte/store';
import data from '$lib/store/data.json';
import Fuse from 'fuse.js';

let recipes = data as unknown as Recipe[];

const fuseOpts = {
	includeScore: true,
	minMatchCharLength: 3,
	ignoreLocation: true,
	threshold: 0,
	keys: ['displayName', 'description']
};

const fuse = new Fuse(recipes, fuseOpts);

class LangStore {
	get allRecipes(): Readable<Recipe[]> {
		return readable(null, function start(set) {
			set(recipes);
			return function stop() {};
		});
	}
}

export const store = new LangStore();
export const textSearch = writable('');
export const selectedLanguage = writable(Languages.none);
export const selectedSignal = writable(Signals.none);
export const selectedRecipeId = writable(Recipes.none.id);

// For drop-downs
export const allLanguages: Readable<SignalDropDown[]> = readable(Languages.all);
export const allSignals: Readable<SignalDropDown[]> = readable(Signals.all);

export const filteredSamples: Readable<Recipe[]> = derived(
	[store.allRecipes, textSearch, selectedLanguage, selectedSignal],
	([$store, $textSearch, $selectedLanguage, $selectedSignal]) => {
		// reset the query params when the selected language/signal changes
		// it's added again when the user selects a sample
		// clearQueryParams();

		let recipes: Recipe[] = $store;

		if (
			!$textSearch &&
			$selectedLanguage.id === Languages.none.id &&
			$selectedSignal.id === Signals.none.id
		) {
			return recipes.sort((a: Recipe, b: Recipe) => a.languageId.localeCompare(b.languageId));
		}

		// if there's any text, filter for it first.
		if ($textSearch) {
			recipes = fuse.search($textSearch).map((r) => r.item);
		}

		if ($selectedLanguage.id !== Languages.none.id) {
			recipes = recipes.filter((r: Recipe) => r.languageId === $selectedLanguage.id);
		}

		// Only signal filter selected
		if ($selectedSignal.id !== Signals.none.id) {
			recipes = recipes.filter((r: Recipe) => r.signal === $selectedSignal.id);
			return recipes;
		}

		return recipes;
	}
);

export const selectedRecipe: Readable<Recipe> = derived(
	[store.allRecipes, selectedRecipeId],
	([$store, $selectedRecipeId]) => {
		if ($selectedRecipeId === Recipes.none.id) {
			return Recipes.none;
		}

		const recipe = $store.find((r: Recipe) => r.id === $selectedRecipeId);
		if (!recipe) {
			return Recipes.none;
		}

		return recipe;
	}
);

export function getLanguage(languageId: string) {
	let l = Languages.all.find((l) => l.id === languageId)
	return l;
}

export function resetSearch() {
	textSearch.set(null);
	selectedLanguage.set(Languages.none);
	selectedSignal.set(Signals.none);
	selectedRecipeId.set(Recipes.none.id);
}

export function setFromUrl(recipeId?: string) {
	if (recipeId) {
		selectedRecipeId.set(recipeId);
		return;
	} else {
		selectedRecipeId.set(Recipes.none.id);
	}
}
