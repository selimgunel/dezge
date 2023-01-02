import { createImmerReducer } from 'ngrx-immer/store';
import { BackendAvailabiltyActions } from '../actions';
import { on } from '@ngrx/store';

export interface BackendAvailablityState {
  available: string | null;
}

const initialState: BackendAvailablityState = {
  available: null,
};

export const availabalityReducer = createImmerReducer(
  initialState,
  on(BackendAvailabiltyActions.availableSuccess, (state, {response}) => {
    state.available = response
    return state;
  }),
  on(BackendAvailabiltyActions.availableFail, (state, action) => {
    return state;
  })
);
