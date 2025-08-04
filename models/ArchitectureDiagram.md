+-----------------------------------------------------+
|          Deep Generative Architecture               |
|                                                     |
|  +---------------------------+                      |
|  |       VAE Module          |                      |
|  |  +-----------+            |                      |
|  |  |  Encoder  |            |                      |
|  |  +-----------+            |                      |
|  |         |                 |                      |
|  |   Latent Space            |                      |
|  |         |                 |                      |
|  |  +-----------+            |                      |
|  |  |  Decoder  |            |                      |
|  |  +-----------+            |                      |
|  +---------------------------+                      |
|               |                                     |
|               v                                     |
|  +-------------------------------+                  |
|  | RNA/Protein Transformer Model |                  |
|  |  (Sequence Generation)        |                  |
|  +-------------------------------+                  |
|                                                     |
|  Input Data -> VAE -> Transformer -> Output         |
+-----------------------------------------------------+


flowchart TD
    A[Input Data] --> B[Encoder (VAE)]
    B --> C[Latent Space]
    C --> D[Decoder (VAE)]
    D --> E[RNA/Protein Transformer]
    E --> F[Output Sequence]
