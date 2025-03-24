import pandas as pd
import seaborn as sns
import matplotlib.pyplot as plt

# Load the data
gene_references = pd.read_csv('gene_references.tsv', sep='\t')
protein_info = pd.read_csv('protein_info.tsv', sep='\t')
pathway_maps = pd.read_csv('pathway_maps.tsv', sep='\t')
functional_insights = pd.read_csv('functional_insights.tsv', sep='\t')

# Example: Generate a heatmap of pathway maps
plt.figure(figsize=(10, 8))
heatmap_data = pathway_maps.pivot("Gene", "Pathway_Name", "Pathway_ID")
sns.heatmap(heatmap_data, annot=True, cmap="YlGnBu")
plt.title("Pathway Maps Heatmap")
plt.savefig("pathway_maps_heatmap.png")

# Example: Generate a volcano plot of functional insights
plt.figure(figsize=(10, 8))
functional_insights['-log10(p-value)'] = -np.log10(functional_insights['Evidence'])
sns.scatterplot(data=functional_insights, x='Fold Change', y='-log10(p-value)', hue='Gene', style='Function_Type')
plt.title("Volcano Plot of Functional Insights")
plt.savefig("functional_insights_volcano_plot.png")

print("Heatmaps and volcano plots generated successfully.")
