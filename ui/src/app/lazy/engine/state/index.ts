import {
  createFeatureSelector,
  createSelector,
} from '@ngrx/store';
import * as fromRoot from '../../../static/store/reducers';
import * as fromEngine from './engine.reducer';

export const engineFeatureKey = 'engine';

export interface State extends fromRoot.AppState {
  engines: fromEngine.EngineInfoState;
}

const getContactFeatureState =
  createFeatureSelector<fromEngine.EngineInfoState>('engine');

export const getEngineList = createSelector(
  getContactFeatureState,
  (state) => state.engines
);
