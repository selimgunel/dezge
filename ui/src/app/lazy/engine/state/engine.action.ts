import { createAction, props } from '@ngrx/store';
import { EngineInfo } from 'src/app/static/models';

export const list = createAction(
  '[Engine List] List'
);

export const listSuccess = createAction(
  '[Engine List] List Success',
  props<{engines: EngineInfo[]}>()
);

export const listFail = createAction(
  '[Engine List] List Fail',
  props<{error: string}>()
);