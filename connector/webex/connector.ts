/*
 * Copyright © 2017. TIBCO Software Inc.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
/// 

import {Injectable} from "@angular/core";
import {WiContrib, WiServiceHandlerContribution, AUTHENTICATION_TYPE} from "wi-studio/app/contrib/wi-contrib";
import {IConnectorContribution, IFieldDefinition, IActionResult, ActionResult} from "wi-studio/common/models/contrib";
import {Observable} from "rxjs/Observable";
import {IValidationResult, ValidationResult, ValidationError} from "wi-studio/common/models/validation";

@WiContrib({})
@Injectable()
export class TibcoWebexConnectorContribution extends WiServiceHandlerContribution {
    constructor() {
        super();
    }

   
    value = (fieldName: string, context: IConnectorContribution): Observable<any> | any => {
        return null;
    }
 
    validate = (name: string, context: IConnectorContribution): Observable<IValidationResult> | IValidationResult => {
      if( name === "Connect") {
         let ConsumerKey : IFieldDefinition;
         let ConsumerSecret: IFieldDefinition;
        
         
         for (let configuration of context.settings) {
    		if( configuration.name ===  "ConsumerKey") {
         ConsumerId = configuration
    		} else if( configuration.name ===  "ConsumerSecret") {
         ConsumerSecret = configuration
    		} 
		 }
     
         

         if ConsumerKey.value && ConsumerSecret.value) {
             
           return ValidationResult.newValidationResult().setReadOnly(false)
         }
          else {
            return ValidationResult.newValidationResult().setReadOnly(true)
        }
      }
       return null;
    }

    action = (actionName: string, context: IConnectorContribution): Observable<IActionResult> | IActionResult => {
       if( actionName == "Connect") {
          return Observable.create(observer => {
         	let ConsumerKey: IFieldDefinition;
         	let ConsumerSecret: IFieldDefinition;
         
         
         	for (let configuration of context.settings) {
    			if( configuration.name ===  "ConsumerKey") {
				 ConsumerKey = configuration;
    			} else if( configuration.name ===  "ConsumerSecret") {
				 ConsumerSecret = configuration;
    			} 
		 	}
          
             let actionResult = {
                context: context,
                authType: AUTHENTICATION_TYPE.BASIC,
                authData: {},
                
        }
      

        observer.next(ActionResult.newActionResult().setSuccess(true).setResult(actionResult));
      
      });     

         
}
       
       return null;
    }
}