/* tslint:disable */
/* eslint-disable */
/**
 * zrok
 * zrok client access
 *
 * The version of the OpenAPI document: 0.3.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface UpdateFrontendRequest
 */
export interface UpdateFrontendRequest {
    /**
     * 
     * @type {string}
     * @memberof UpdateFrontendRequest
     */
    frontendToken?: string;
    /**
     * 
     * @type {string}
     * @memberof UpdateFrontendRequest
     */
    publicName?: string;
    /**
     * 
     * @type {string}
     * @memberof UpdateFrontendRequest
     */
    urlTemplate?: string;
}

/**
 * Check if a given object implements the UpdateFrontendRequest interface.
 */
export function instanceOfUpdateFrontendRequest(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function UpdateFrontendRequestFromJSON(json: any): UpdateFrontendRequest {
    return UpdateFrontendRequestFromJSONTyped(json, false);
}

export function UpdateFrontendRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UpdateFrontendRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'frontendToken': !exists(json, 'frontendToken') ? undefined : json['frontendToken'],
        'publicName': !exists(json, 'publicName') ? undefined : json['publicName'],
        'urlTemplate': !exists(json, 'urlTemplate') ? undefined : json['urlTemplate'],
    };
}

export function UpdateFrontendRequestToJSON(value?: UpdateFrontendRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'frontendToken': value.frontendToken,
        'publicName': value.publicName,
        'urlTemplate': value.urlTemplate,
    };
}

