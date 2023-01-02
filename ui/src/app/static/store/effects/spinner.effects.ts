import { Injectable } from '@angular/core';
import { createEffect } from '@ngrx/effects';
import { select, Store } from '@ngrx/store';
import { tap } from 'rxjs/operators';
import { SpinnerOverlayService } from '../../services/spinner-overlay.service';
import * as fromStore from '../reducers';


@Injectable()
export class SpinnerEffects {

    constructor(
      private store$: Store<fromStore.AppState>,
      private spinner: SpinnerOverlayService) {}

    handleSpinner$ = createEffect(
      () =>
        this.store$.pipe(
          select(fromStore.getSpinnerShow),
          tap( isShow =>
            isShow ? this.spinner.show() : this.spinner.hide()
        )),
        { dispatch: false }

    );
}
