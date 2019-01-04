// WARNING: This file has automatically been generated on Fri, 04 Jan 2019 16:12:37 CST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package faiss_go

/*
#cgo darwin LDFLAGS: -Lc_api -lfaiss_c
#cgo darwin LDFLAGS: -L.. -lfaiss
#cgo linux LDFLAGS: -Lc_api -lfaiss_c
#cgo linux LDFLAGS: -L.. -lfaiss
#cgo darwin LDFLAGS: -Wl,-undefined,dynamic_lookup
#cgo darwin LDFLAGS: -lblas -llapack
#cgo linux LDFLAGS: -L/usr/local/lib -lstdc++ -pthread -lcrypt -lm -fopenmp -lblas -llapack
#include "c_api/AutoTune_c.h"
#include "c_api/AuxIndexStructures_c.h"
#include "c_api/Clustering_c.h"
#include "c_api/Index_c.h"
#include "c_api/IndexFlat_c.h"
#include "c_api/IndexIVF_c.h"
#include "c_api/IndexLSH_c.h"
#include "c_api/MetaIndexes_c.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

// Faissmetrictype as declared in c_api/Index_c.h:32
type Faissmetrictype int32

// Faissmetrictype enumeration from c_api/Index_c.h:32
const (
	MetricInnerProduct Faissmetrictype = iota
	MetricL2           Faissmetrictype = 1
)
