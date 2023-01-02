import { Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { map, tap } from 'rxjs/operators';
import { AppearanceColor, SnackBarInterface } from '../../models';
import { ErrorActions, SnackBarActions } from '../actions';





@Injectable()
export class ErrorEffects {

    constructor(private readonly actions$: Actions) {
    }

    handleError$ = createEffect(
      () =>
        this.actions$.pipe(
          ofType(ErrorActions.errorMessage),
          map(action => action.errorMsg),
          tap(errorMsg => console.error('Got error:', errorMsg)),
          map(errorMsg => {

            const msg: SnackBarInterface = {
              message: errorMsg ? errorMsg : '',
              color: AppearanceColor.Error
          };

            return SnackBarActions.open({payload: msg});
          } )
        )

    );

}
