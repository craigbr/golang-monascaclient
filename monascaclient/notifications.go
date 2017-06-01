// Copyright 2017 Hewlett Packard Enterprise Development LP
//
//    Licensed under the Apache License, Version 2.0 (the "License"); you may
//    not use this file except in compliance with the License. You may obtain
//    a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//    License for the specific language governing permissions and limitations
//    under the License.

package monascaclient

import (
	"github.com/monasca/golang-monascaclient/monascaclient/models"
	"encoding/json"
)

func GetNotificationMethods() (*models.NotificationMethodsResponse, error){
	return monClient.GetNotificationMethods()
}

func (p *Client) GetNotificationMethods() (*models.NotificationMethodsResponse, error) {
	monascaURL, URLerr := p.createMonascaAPIURL("v2.0/notification-methods",  map[string]string{}, nil)
	if URLerr != nil {
		return nil, URLerr
	}

	body, monascaErr := p.callMonasca(monascaURL)
	if monascaErr != nil {
		return nil, monascaErr
	}

	notificationMethodsResponse := models.NotificationMethodsResponse{}
	err := json.Unmarshal(body, &notificationMethodsResponse)
	if err != nil {
		return nil, err
	}

	return &notificationMethodsResponse, nil
}