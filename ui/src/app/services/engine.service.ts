import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Engine } from './engines';
import { Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { EngineInfo } from '../static/models';

@Injectable({
  providedIn: 'root'
})
export class EngineService {
  
  constructor(
    private http: HttpClient
  ) {}
 
  getEngines(): Observable<EngineInfo[]> {
    const url = `${environment.apiUrl}/engines`;
    return this.http.get<EngineInfo[]>(url);
  }
  ping(): Observable<string> {
    const url = `${environment.apiUrl}/ping`
    return this.http.get<string>(url);
  }
}


