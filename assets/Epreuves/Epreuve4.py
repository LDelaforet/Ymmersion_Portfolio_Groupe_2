plusLongLength = 0

plusLongMot = ""

nbreChar = 0

mots = []

texte = "Lorem ipsum dolor sit amet"

mots = texte.split(" ")

for x in mots:
	if len(x) > plusLongLength:
		plusLongLength = len(x)
		plusLongMot = x

print("Le mot fait:", len(mots), "mots")
print("Le texte fait", len(texte), "caractères")
print("Le mot le plus long est", plusLongMot)