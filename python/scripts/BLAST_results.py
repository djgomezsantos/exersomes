from Bio.Blast import NCBIXML

def parse_blast_xml(blast_xml_file):
    with open(blast_xml_file) as result_handle:
        blast_records = NCBIXML.parse(result_handle)
        for blast_record in blast_records:
            for alignment in blast_record.alignments:
                for hsp in alignment.hsps:
                    print(f"Match: {alignment.hit_def}, Score: {hsp.score}, Sequence: {hsp.sbjct}")

# Call the function with your BLAST XML file
parse_blast_xml("your_blast_output.xml")
