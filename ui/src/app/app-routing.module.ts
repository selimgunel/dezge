import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import {DocComponent} from './lazy/doc/doc.component'


const routes: Routes = [
  {path: 'doc', component:DocComponent},
  {
    path: 'engines',
    loadChildren: () => import('./lazy/engine/engine.module').then(m => m.EngineModule)
  },
  {
    path: 'tournaments',
    loadChildren: () => import('./lazy/tournament/tournament.module').then(m => m.TournamentModule)
  },

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
