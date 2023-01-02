import * as EngineActions from './engine.action';
import { EngineInfo } from './../../../static/models/engine';
import { createImmerReducer } from 'ngrx-immer/store';
import { on } from '@ngrx/store';

export interface EngineInfoState {
  engines: EngineInfo[] | null;
  error: string | null;
}

export const initialState: EngineInfoState = {
  engines: null,
  error: null
};



export const engineInfoReducer = createImmerReducer(
	initialState,
	on(EngineActions.listSuccess, (state, action) => {
		state.engines = action.engines;
		return state;
	}),
	on(EngineActions.listFail, (state, action) => {
		state.error = action.error
		return state;
	}),
);