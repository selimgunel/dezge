import { TestBed } from '@angular/core/testing';

import { EngineService } from './engine.service';
// Http testing module and mocking controller
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Engine } from './engines';
import { EngineInfo } from '../static/models';

describe('EngineService', () => {

  let httpClient: HttpClient;
  let httpTestingController: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [ HttpClientTestingModule ]
    });

    // Inject the http service and test controller for each test
    httpClient = TestBed.inject(HttpClient);
    httpTestingController = TestBed.inject(HttpTestingController);
  });

  let service: EngineService;
  const testUrl = `http://localhost:8080`
  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('can test HttpClient.getEngines', () => {
    const testData: EngineInfo[] = [
      {id: 1, name: 'Test Data', path: 'Path', engine_id: 'engine_id', createdAt: Date.now.toString(),  updatededAt: Date.now.toString()}];
  
    // Make an HTTP GET request
    httpClient.get<EngineInfo[]>(testUrl)
      .subscribe(data =>
        // When observable resolves, result should match test data
        expect(data).toEqual(testData)
      );
  
    // The following `expectOne()` will match the request's URL.
    // If no requests or multiple requests matched that URL
    // `expectOne()` would throw.
    const req = httpTestingController.expectOne('/engines');
  
    // Assert that the request is a GET.
    expect(req.request.method).toEqual('GET');
  
    // Respond with mock data, causing Observable to resolve.
    // Subscribe callback asserts that correct data was returned.
    req.flush(testData);
  
    // Finally, assert that there are no outstanding requests.
    httpTestingController.verify();
  });

  afterEach(() => {
    // After every test, assert that there are no more pending requests.
    httpTestingController.verify();
  });
});
