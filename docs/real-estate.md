# Buy My House

This page describes how we leverage AI technologies and `jimi` CLI to analyse real-restate offers.

* Design plan
* Create prompts
* Extract information from interested real estate offers
* Generate analysis report

## Design Plan

```
### Context ###
I want to buy a house in the greater Paris area. I want you to play the role of
a real-estate agent to assist me on real estate analysis. Keep in mind that I
am already an owner of a property and having several years of experience living
in this area. So don't afraid to use technical and specific terms and don't
assume that I am a beginner.

### Instruction ###
Assume that I have the description of the house, downloaded from a real estate
website, such as SeLoger or a local real-estate agency. I want you to assist me
in making decisions for deciding whether this house is a good candidate. Please
make a plan about the analysis. You don't need to perform actual analysis for
now, I just need a framework to iterate on actual content later on.

### Input Data ###
None

### Output Indicator ###
Group the output by different aspects of the purchase as "dimension". Describe
the definition of the dimension. Explain the interest of exploring this
dimension. Explain how can I find related information from the content of a
real estate offer. What additional information is required from me or from
other data sources is required to perform the analysis. Be creative if you
think other info should be included in the output.
```