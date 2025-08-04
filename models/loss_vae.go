import "math"

// ReconstructionLoss computes MSE between original and reconstructed data
func ReconstructionLoss(original, reconstructed []float64) float64 {
    var loss float64
    for i := range original {
        diff := original[i] - reconstructed[i]
        loss += diff * diff
    }
    return loss / float64(len(original))
}

// KLDivergence computes KL divergence between q(z|x) ~ N(mu, var) and p(z) ~ N(0,1)
func KLDivergence(mu, logVar LatentVector) float64 {
    var kl float64
    for i := range mu {
        kl += 1 + logVar[i] - mu[i]*mu[i] - math.Exp(logVar[i])
    }
    return -0.5 * kl
}

// VAELoss combines reconstruction and KL loss
func VAELoss(original, reconstructed []float64, mu, logVar LatentVector, beta float64) float64 {
    rec := ReconstructionLoss(original, reconstructed)
    kl := KLDivergence(mu, logVar)
    return rec + beta*kl // beta=1 for standard VAE, >1 for beta-VAE
}
