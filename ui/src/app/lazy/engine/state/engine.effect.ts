import { Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { catchError, map, switchMap, tap } from 'rxjs/operators';
import { serializeError } from 'serialize-error';
import { EngineInfo } from '../../../static/models/engine';
import * as engineActions from './engine.action';

import { ErrorActions, SpinnerActions } from '../../../static/store/actions';
import { EngineService } from 'src/app/services/engine.service';
import { of } from 'rxjs';


@Injectable()
export class EngineEffects {

    constructor(
        private readonly actions$: Actions,
        private readonly engineService: EngineService
        ) {
    }

    getEngineList$ = createEffect(
      () =>
        this.actions$.pipe(
          ofType(engineActions.list),
         switchMap(()=>
            this.engineService.getEngines().pipe(
                map((engines: EngineInfo[]) => engineActions.listSuccess({engines}) ),
                catchError((err: Error) => of(this.handleError(err)))
            )))
    );

    private handleError(error: Error) {
        const friendlyErrorMessage = serializeError(error).message;
        return ErrorActions.errorMessage({ errorMsg: friendlyErrorMessage });
      }
}
