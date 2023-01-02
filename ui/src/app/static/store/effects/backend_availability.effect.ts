import { Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { catchError, map, switchMap, tap } from 'rxjs/operators';
import { serializeError } from 'serialize-error';
import { EngineService } from 'src/app/services/engine.service';
import { ErrorActions, SnackBarActions } from '../actions';
import * as backendAvailablityActions from './../actions/backend_availability.action';
import {  of } from 'rxjs';

@Injectable()
export class BackendAvailablityEffects {
  constructor(
    private readonly actions$: Actions,
    private readonly engineService: EngineService
  ) {}

  getBackendAvailability$ = createEffect(() =>
    this.actions$.pipe(
      ofType(backendAvailablityActions.available),
      switchMap(() =>
        this.engineService.ping().pipe(
          map((response: string) =>
            backendAvailablityActions.availableSuccess({ response })
          ),
          catchError((err: Error) => of(this.handleError(err)))
        )
      )
    )
  );

  private handleError(error: Error) {
    const friendlyErrorMessage = serializeError(error).message;
    return ErrorActions.errorMessage({ errorMsg: friendlyErrorMessage });
  }
}
