// https://github.com/tomastrajan/angular-ngrx-material-starter
// https://medium.com/@aleixsuau/error-handling-angular-859d529fa53a
// https://medium.com/angular-in-depth/expecting-the-unexpected-best-practices-for-error-handling-in-angular-21c3662ef9e4
// https://www.tektutorialshub.com/angular/error-handling-in-angular-applications/

import { HttpErrorResponse } from '@angular/common/http';
import { ErrorHandler, Injectable, Injector } from '@angular/core';
import { serializeError } from 'serialize-error';
import { AppearanceColor } from '../models';
import { SnackBarService } from '../services/snack-bar.service';


@Injectable({ providedIn: 'root' })
export class AppErrorHandler extends ErrorHandler {

  // Error handling is important and it is loaded first.
  // Because of this we should manually inject the services with Injector.
  constructor(private injector: Injector) {
    super();
   }

  override handleError(error: Error | HttpErrorResponse) {

    const notifier = this.injector.get(SnackBarService);
    let message: string | undefined;

    if (error instanceof HttpErrorResponse) {
      // Server Error
      message = error.message;
    } else {
      // Client Error
      message = serializeError(error).message;
    }
    
    message = message ? message : '';
    notifier.show({
      message,
      color: AppearanceColor.Error
    });

    // continue with the default global error handler
    super.handleError(error);

  }

}
