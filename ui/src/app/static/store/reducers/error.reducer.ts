import { createReducer, on } from '@ngrx/store';
import { ErrorActions } from '../actions';

export interface ErrorState {
  error: string | undefined;
}

const initialState: ErrorState = {
  error: undefined
};

export const reducer = createReducer(
  initialState,
  on(ErrorActions.errorMessage, (state, { errorMsg }) => ({
    ...state,
     error: errorMsg
    })),
);

export const getError = (state: ErrorState) => state.error;

