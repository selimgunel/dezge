import { NgModule, isDevMode, ErrorHandler } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { StoreModule } from '@ngrx/store';
import { StoreDevtoolsModule } from '@ngrx/store-devtools';
import { EffectsModule } from '@ngrx/effects';
import { ErrorEffects } from './static/store/effects/error.effects';
import { environment } from 'src/environments/environment';
import { MaterialModule } from './material';
import { ToolbarComponent } from './static/ui/components/toolbar/toolbar.component';
import { DocComponent } from './lazy/doc/doc.component';
import { ROOT_REDUCERS, metaReducers } from './static/store/reducers/index';
import { HttpErrorInterceptor } from './static/services/http-error.interceptor';
import { AppErrorHandler } from './static/services/app-error-handler.service';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { SnackBarComponent } from './static/ui/components/snackbar/snack-bar.component';

const EFFECTS = [
  ErrorEffects
]

@NgModule({
  declarations: [
    AppComponent,
    DocComponent,
    ToolbarComponent,
    SnackBarComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    HttpClientModule,
    MaterialModule,
    StoreModule.forRoot(ROOT_REDUCERS, { metaReducers }),
    !environment.production ? StoreDevtoolsModule.instrument({name: 'Starter App'}) : [] ,
    EffectsModule.forRoot(EFFECTS),

  ],
  providers: [
    // register "the GlobalErrorHandler provider
    { provide: ErrorHandler, useClass: AppErrorHandler },
    { provide: HTTP_INTERCEPTORS, useClass: HttpErrorInterceptor, multi: true },
  ],  bootstrap: [AppComponent],

})
export class AppModule { }
