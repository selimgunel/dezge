
import { InjectionToken } from '@angular/core';
import {
  Action, ActionReducer, ActionReducerMap, createFeatureSelector, createSelector, MetaReducer
} from '@ngrx/store';
import { environment } from '../../../../environments/environment';
import * as fromError from './error.reducer';
import * as fromSpinner from './spinner.reducer';
import * as fromEngine from './../../../lazy/engine/state/engine.reducer';

// nice moment here
// here is our root state, which also includes the route state
export interface AppState {
  spinner: fromSpinner.SpinnerState;
  error: fromError.ErrorState;
  engines: fromEngine.EngineInfoState
}

/**
 * Our state is composed of a map of action reducer functions.
 * These reducer functions are called with each dispatched action
 * and the current or initial state and return a new immutable state.
 */
export const ROOT_REDUCERS =
  new InjectionToken<ActionReducerMap<AppState, Action>>('Root reducers token', {
    factory: () => ({
      spinner: fromSpinner.reducer,
      error: fromError.reducer,
      engines: fromEngine.engineInfoReducer
    }),
  });

// console.log all actions
export function logger(reducer: ActionReducer<AppState>): ActionReducer<AppState> {
  return (state, action) => {
    const result = reducer(state, action);
    console.groupCollapsed(action.type);
    console.log('prev state', state);
    console.log('action', action);
    console.log('next state', result);
    console.groupEnd();

    return result;
  };
}

/**
 * By default, @ngrx/store uses combineReducers with the reducer map to compose
 * the root meta-reducer. To add more meta-reducers, provide an array of meta-reducers
 * that will be composed to form the root meta-reducer.
 */
export const metaReducers: MetaReducer<AppState>[] = !environment.production
  ? [logger]
  : [];


export const selectSpinnerState = createFeatureSelector<fromSpinner.SpinnerState>(
  'spinner'
);
export const getSpinnerShow = createSelector(
  selectSpinnerState,
  fromSpinner.getSpinnerShow
);


export const selectErrorState = createFeatureSelector<fromError.ErrorState>(
  'error'
);

export const getError = createSelector(
  selectErrorState,
  fromError.getError
);
