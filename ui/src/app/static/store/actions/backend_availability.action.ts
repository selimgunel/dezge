import { createAction, props } from '@ngrx/store';
import { EngineInfo } from 'src/app/static/models';

export const available = createAction(
  '[Main Page] Available'
);

export const availableSuccess = createAction(
  '[Main Page] Available Success',
  props<{response: string}>()
);

export const availableFail = createAction(
  '[Main Page] Available Fail',
  props<{error: string}>()
);