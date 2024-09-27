package blst

// Pop used in the BLS signature scheme.
type ProofOfPossesion struct {
	s *blstSignature
}

// Marshal a signature into a LittleEndian byte slice.
func (p *ProofOfPossesion) Marshal() []byte {
	return p.s.Compress()
}

func (p *ProofOfPossesion) Equals(p2 Pop) bool {
	return p.s.Equals(p2.(*ProofOfPossesion).s)
}

// Copy returns a full deep copy of a signature.
func (p *ProofOfPossesion) Copy() Pop {
	sign := *p.s
	return &ProofOfPossesion{s: &sign}
}
