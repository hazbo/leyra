// leyra v0.0.1-dev
//
// (c) Ground Six 2015
//
// @package leyra
// @version 0.0.1-dev
//
// @author Harry Lawrence <http://github.com/hazbo>
//
// License: MIT
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package bootstrap

import (
	"log"

	"gopkg.in/leyra/godotenv.v1"
)

// SetEnv tries to read the .env file in the root of this project and load in
// the needed env vars.
func SetEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
