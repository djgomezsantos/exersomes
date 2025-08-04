package models

import (
    "math"
    "math/rand"
)

// LatentVector represents the z latent space of the VAE
type LatentVector []float64

// VAEEncoder encodes inputs to a latent space z
type VAEEncoder struct {
    Mu    LatentVector // Mean vector
    LogVar LatentVector // Log-variance vector
    Z     LatentVector // Latent variable
}

// NewVAEEncoder creates a new encoder with specified latent size
func NewVAEEncoder(latentDim int) *VAEEncoder {
    return &VAEEncoder{
        Mu:    make(LatentVector, latentDim),
        LogVar: make(LatentVector, latentDim),
        Z:     make(LatentVector, latentDim),
    }
}

// Reparameterize samples z using the reparameterization trick
func (e *VAEEncoder) Reparameterize() {
    for i := range e.Z {
        std := math.Exp(0.5 * e.LogVar[i])
        epsilon := rand.NormFloat64()
        e.Z[i] = e.Mu[i] + std*epsilon
    }
}

// Example: encode some input (pseudo-code)
// func (e *VAEEncoder) Encode(input []float64) {
//     // Compute e.Mu and e.LogVar from input...
//     // Call e.Reparameterize() to obtain z
// }
