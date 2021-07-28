/* Copyright 2010 The Go Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.
*
* Source modified by Andrew Rodman to work with the customized
* PSOBB Blowfish implementation. Work based off of the encryption
* library written by Fuzziqer Software.
 */

package encryption

func initCipher(c *blowfishCipher) {
	copy(c.p[0:], p[0:])
	copy(c.s0[0:], s0[0:])
	copy(c.s1[0:], s1[0:])
	copy(c.s2[0:], s2[0:])
	copy(c.s3[0:], s3[0:])
}

func expandKey(key []byte, c *blowfishCipher) {
	keyCopy := make([]byte, len(key))
	copy(keyCopy, key)
	// PSO applies a fixed salt to its encryption keys.
	for i := 0; i < 48; i += 3 {
		keyCopy[i] ^= 0x19
		keyCopy[i+1] ^= 0x16
		keyCopy[i+2] ^= 0x18
	}

	// PSO does some de-obfuscation of the P table. Specifically, for each P table
	// entry it reverses the 2 least significant bytes, shifts them to the upper
	// 16 bits and XORs the former most significant bytes into the lower half.
	for i := 0; i < 18; i++ {
		entry := p[i]
		upper := entry & 0xffff
		upper = ((upper & 0xff) << 8) + (upper >> 8)
		lower := entry & 0xffff0000
		lower ^= (upper << 16)
		c.p[i] = upper | lower

	}

	// XOR each entry of the P table with the salted key. Each key
	// is combined with its adjacent entry for a total of 4 bytes,
	// wrapping around at the end of the key.
	j := 0
	for i := 0; i < 18; i++ {
		var d uint32
		for k := 0; k < 4; k++ {
			d = d<<8 | uint32(keyCopy[j])
			j++
			if j >= len(keyCopy) {
				j = 0
			}
		}
		c.p[i] ^= d
	}

	var l, r uint32
	for i := 0; i < 18; i += 2 {
		l, r = encryptPBlock(l, r, c)
		c.p[i], c.p[i+1] = l, r
	}
	for i := 0; i < 256; i += 2 {
		l, r = encryptSBlock(l, r, c)
		c.s0[i], c.s0[i+1] = l, r
	}
	for i := 0; i < 256; i += 2 {
		l, r = encryptSBlock(l, r, c)
		c.s1[i], c.s1[i+1] = l, r
	}
	for i := 0; i < 256; i += 2 {
		l, r = encryptSBlock(l, r, c)
		c.s2[i], c.s2[i+1] = l, r
	}
	for i := 0; i < 256; i += 2 {
		l, r = encryptSBlock(l, r, c)
		c.s3[i], c.s3[i+1] = l, r
	}
}

// PSO uses the standard Blowfish rounds for P table scheduling.
func encryptPBlock(l, r uint32, c *blowfishCipher) (uint32, uint32) {
	xl, xr := l, r
	xl ^= c.p[0]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[1]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[2]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[3]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[4]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[5]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[6]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[7]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[8]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[9]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[10]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[11]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[12]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[13]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[14]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[15]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[16]
	xr ^= c.p[17]
	return xr, xl
}

// Mostly the same as standard Blowfish, but with some customization at the end.
func encryptSBlock(l, r uint32, c *blowfishCipher) (uint32, uint32) {
	xl, xr := l, r
	xl ^= c.p[0]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[1]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[2]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[3]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[4]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[5]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[6]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[7]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[8]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[9]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[10]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[11]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[12]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[13]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[14]

	tmp := (((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)]) ^ c.p[15]
	tmp ^= xr

	xr = (((c.s0[byte(tmp>>24)] + c.s1[byte(tmp>>16)]) ^ c.s2[byte(tmp>>8)]) + c.s3[byte(tmp)]) ^ c.p[16]
	xr ^= xl

	xl = c.p[17]
	xl ^= tmp

	return xl, xr
}

func encryptData(l, r uint32, c *blowfishCipher) (uint32, uint32) {
	xl, xr := l, r
	xl ^= c.p[0]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[1]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[2]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[3]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[4]
	xr ^= c.p[5]
	return xr, xl
}

func decryptData(l, r uint32, c *blowfishCipher) (uint32, uint32) {
	xl, xr := l, r
	xl ^= c.p[5]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[4]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[3]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[2]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[1]
	xr ^= c.p[0]
	return xr, xl
}

var s0 = [256]uint32{
	0x5cc73fd6, 0x19572a8e, 0x1ead320e, 0x29913b33, 0x05c06104, 0xc5a1316e,
	0x456d82a7, 0x5a987789, 0xbfdcaa97, 0x23094413, 0x70b100f7, 0xef18f524, 0x9b2632b1, 0x1a7aa450,
	0x36355519, 0x1a8fc2ad, 0xe13d6a17, 0xc74af6af, 0xb771fc73, 0x8c332a8c, 0x13792c10, 0xa707616f,
	0x69d18ce4, 0x4bb744c2, 0x74da584b, 0xc2186564, 0xbef96bfd, 0x9ff42f9e, 0xf6334290, 0x74249103,
	0xc0d5cccc, 0xacc295f2, 0x7bb4d473, 0xba4753a2, 0x46806e9f, 0x7f8eb321, 0x16803fe6, 0x891fd2bc,
	0xe7218373, 0x82cdf207, 0x819879e3, 0xb0cde742, 0xa9160843, 0x14336f73, 0xfc0b51a2, 0x2fd13817,
	0x231c4134, 0xbea851c9, 0x1b8b5dbf, 0xb225a875, 0x6bcc7ec4, 0xcc5c66ef, 0xb5c06e89, 0xdc479976,
	0x1b984ed0, 0x35bd70c1, 0x8ec73f26, 0xc85ed3fb, 0x93a2cff3, 0xc0c1889f, 0x74c405a6, 0xb4ba3842,
	0x62a89a52, 0x850373d1, 0xa8ad015e, 0x4087946a, 0x1e81c985, 0xe0278fee, 0x6d38ec05, 0xbf4e158a,
	0x63e32bdd, 0x17578163, 0x9861c874, 0x535ed4cf, 0xe0674a4a, 0xa2233b6c, 0x523574e3, 0x35d19568,
	0x0247aff9, 0xed2bf2a5, 0x1a404cc6, 0x5700a52c, 0x3f5847fc, 0x9139f9fc, 0x8721985a, 0x17a0493b,
	0xf333a0d4, 0x489411ff, 0xd92ef4db, 0x1e50c960, 0x833757f6, 0xbbc00c1b, 0x01558f29, 0x68058035,
	0xdb7d2645, 0x5d11a667, 0xaee02660, 0x3f5b5474, 0xe3cf9e79, 0x8e0e6574, 0x2e4cec6f, 0xb900ebb0,
	0x30f703c1, 0xe73ecee4, 0x907d69cb, 0xb785648d, 0xee57bbe1, 0xa0862cb0, 0xe942e1c5, 0x2c7d0221,
	0xdf5445f7, 0xd8c8ac9f, 0x22f05641, 0x3e295eac, 0x1e138faa, 0x3598f64d, 0xda199769, 0xf157c46d,
	0x7ca171c5, 0x94301db9, 0xfdc90d52, 0x387128a3, 0x41d7c806, 0xd3190dab, 0x3acd7a85, 0xe83ebbe3,
	0x14322c57, 0x26845b42, 0xb2cd49cb, 0xe4d22b24, 0x23c11989, 0xe4fcd996, 0x0fc3ad3d, 0xe17a680c,
	0xf4f0f8d8, 0x72350d14, 0x4c747633, 0xc9633b10, 0xfec3618b, 0xfde8dd1c, 0x9369edf4, 0xc8aeced7,
	0xe7160549, 0x75bd584c, 0xf0451846, 0xcefb421c, 0x50fc8705, 0xd67643ae, 0x970afde8, 0x09f8deba,
	0x6e82eaad, 0x80ceb947, 0x51afe307, 0x727b3f2f, 0xb22b287b, 0xf077f03a, 0x4b670178, 0x1f942dde,
	0x37afeaff, 0xe569cde3, 0xb78dd11d, 0x6e8307d1, 0x95ce57c6, 0xc0e34476, 0x2ca562d1, 0x6373d161,
	0x2e549898, 0xc6f47ec6, 0x4a2a6be4, 0x6898dd70, 0xff954a7c, 0x8f033cd0, 0xcd64c8e8, 0x3c0a7d7b,
	0xa3057d95, 0xecd438e0, 0xc111363a, 0xb94fd214, 0x7f224dfe, 0xf042a491, 0x9f1489fc, 0x75e73dc9,
	0x1ea04f71, 0xa38f2685, 0x8ba7af61, 0x8dbf33df, 0x4eacd05d, 0x3cef9b0e, 0x9604fe9f, 0xb65d9990,
	0xbbcb14ba, 0x06fc3a41, 0xe15376de, 0x97d9bc59, 0x8318618a, 0x2db10c0c, 0x3736fc1f, 0x6e8136d8,
	0x7e470db5, 0xc60daed2, 0x5a19532f, 0x98094aa8, 0xe830fea2, 0x126a0685, 0x2b76b98f, 0xa378f291,
	0xd36fb474, 0xa3849120, 0x7868242a, 0x87743ea2, 0x1d74914f, 0xb341998c, 0xd5b45b60, 0xcd97dd2e,
	0x9cef94c3, 0x907d0c7b, 0xaa967285, 0x2c0c2b35, 0x852d480b, 0xafb7455c, 0x0fb40a91, 0xde019ac6,
	0xf285aa86, 0xb5af214b, 0x94e3a9d8, 0x61cc82de, 0x592ce330, 0x24943eef, 0xc689113f, 0x68a7fe73,
	0xce85cae0, 0x9477d5b7, 0x7ee161b8, 0x1c4f6b1b, 0xad1073f1, 0xfba9fff8, 0x11a5ce22, 0x19be7af3,
	0x8646d47a, 0xdd92e45b, 0xa5b089c8, 0x05db18a7, 0xd915fb67, 0xae545e52, 0x738b8333, 0xe351e074,
	0xd846f324, 0x4c4c85ae, 0x1f705eaf, 0x3c65970c, 0xb540a652, 0x08355576, 0x88fd52f2, 0x1176fa93,
	0x04d2406a, 0xa53e17c7,
}

var s1 = [256]uint32{
	0xc5fb6441, 0xd36fc212, 0x5c5ac0c9, 0xe2c932c2, 0xd22a7467, 0xad1d4b06,
	0xdc30354a, 0x09f640ea, 0x1b063309, 0x0777b7a2, 0xe30f2845, 0xb16ed5e6, 0x897b6abf, 0x1e2ec223,
	0xcfb0ac5c, 0x0297f232, 0x7f56f89d, 0xa3f50491, 0x7c847191, 0x61d4b903, 0x25ee2690, 0x58f77a26,
	0xc2d527fe, 0x8123afbe, 0x7dff42e6, 0x9572104b, 0x15d8e9f6, 0x23f908c8, 0x1156a4dc, 0xf8816e83,
	0xaea972ec, 0x9095ecfb, 0xfdd7afad, 0xaa156f86, 0x3306c3ad, 0x5b21343d, 0x13d0f0d9, 0xa9098abf,
	0x522944f1, 0x76d2a256, 0xe259a0b5, 0x4675d80d, 0x8b3dfc79, 0xb9a76f83, 0xf168cd53, 0x0609a55b,
	0x98e96452, 0xb17832d9, 0x8a90cbc9, 0xc0229573, 0x17266917, 0x20055f24, 0xaaf79b0d, 0xe0d393eb,
	0x282c0b07, 0x63af3bbe, 0x9fd9ae8f, 0xa0325e5c, 0x759b22ac, 0xabb02882, 0xaa56e55c, 0xa302aa9e,
	0x95e40019, 0x1f41e3c9, 0x164b605d, 0x30cd6081, 0xf46f6677, 0x66fddbb7, 0xae738ace, 0x64a9b3ff,
	0x76cb795f, 0x8671b0e4, 0x946fdf07, 0x0f0712dc, 0x14be281a, 0xee01e411, 0x5473c49f, 0xdd572435,
	0x6183d89b, 0xd8946913, 0xcca66ff9, 0x39d5a9be, 0x3b1a7d18, 0xa72b5d96, 0x111e8e30, 0xdab26740,
	0x3f64b3de, 0xd1695e1a, 0x33a19648, 0x31dc630a, 0xf5f35694, 0xb91ed674, 0x06fe9043, 0xbe9e4e5b,
	0xda426aab, 0x535055ec, 0x0d2b265e, 0xef43b103, 0xb7edf4b1, 0xaa2618f5, 0xa3d00018, 0xefa242cc,
	0x49d47f55, 0x562677c2, 0x7d41eeda, 0x40bf3aa5, 0x135b8eea, 0x5dbed1da, 0x99bc688a, 0xfe073b61,
	0x34e62a8e, 0x5125d336, 0xdc70a9a6, 0x292c52b4, 0x2c7e2f60, 0x04647f1f, 0x8a1989c4, 0xeba69244,
	0xa54a3897, 0xfac0d4d0, 0xad47205b, 0xf794c013, 0xcd3c0a23, 0xba9671ac, 0x8d1eaea6, 0x0de2e83e,
	0x9fbee730, 0xfa0684a3, 0x42d96104, 0x0e97ce42, 0x698374a6, 0xef7d8288, 0xf590de72, 0x6899f987,
	0x1bfd58ec, 0x38b4b274, 0x088a50ae, 0xae2113b0, 0xe64cf295, 0xbb67f9be, 0xdfa77bc0, 0x598481ea,
	0x13e267b8, 0xa7eb1033, 0x7ca6ddda, 0x4a836ced, 0xbf89c618, 0xbfbe9dae, 0xa44fe33a, 0xa0be3198,
	0xed12af84, 0x20976be3, 0x4754aa5b, 0x72930c88, 0xb68d8550, 0x532558e9, 0x230f5f40, 0xc0bd9035,
	0x672f3482, 0xa89a61bf, 0x4aa288dc, 0x2045c67a, 0xc59b9ae6, 0xa0337df9, 0xe1857270, 0xfd3dff5d,
	0xe301ec12, 0x50ffae66, 0x89dfe89c, 0x768c6e14, 0x0aa10d87, 0xfe2feef1, 0x61b3a2ee, 0xd5a31e6f,
	0x7789b9f2, 0x0ff5c3b1, 0x29f1c194, 0x77d011b0, 0xecc10b84, 0x6f931750, 0x70b62a8f, 0xbb83cefe,
	0x5f497eb2, 0xf17666d6, 0x5d785704, 0x865c980b, 0xf0249ec2, 0xaae844db, 0x4cd28e52, 0xda93ade9,
	0xd966908c, 0xa4b9fddc, 0x1fae7671, 0x96513d07, 0x98f07cb6, 0x7c13b222, 0x1f05ffe7, 0xff903b48,
	0xc8d0dbbb, 0xa6e52eb5, 0x7d7bc10a, 0xafe0d2f7, 0x01b79cc8, 0x578225f9, 0xe40c41b3, 0xb5c7e26a,
	0xa46286ef, 0x7b138d12, 0x432661b3, 0xc9c8124e, 0xe4be379b, 0x34aee10d, 0x59aff4cb, 0xdad26c27,
	0x5c9561b8, 0x4d6b0452, 0x10955f82, 0x8aad8718, 0x4aaf2843, 0xb94c51f7, 0x756ff181, 0xe701f22e,
	0xa70427ee, 0x52654509, 0x2e4c3cab, 0x33e7af57, 0xccdc8f42, 0xb8b3ca13, 0x9122c3f3, 0xf074441a,
	0x48e1c890, 0xae102653, 0xc977a7f2, 0x2fe76749, 0x754513c2, 0xa2a86df9, 0x7312f6b7, 0xcca4e105,
	0xacfb96cd, 0xa0a9a9b2, 0x237faf6d, 0x45b7eb4d, 0x0c3e5872, 0x460c5991, 0x97248330, 0xa47541b2,
	0xbf76d53b, 0x6c6c782b, 0x38a76a50, 0x712e9fec, 0xe7071507, 0x0e4202b2, 0x95a4154e, 0x62f6da87,
	0x3dfd5418, 0xd7ab33f5,
}

var s2 = [256]uint32{
	0x8e13062c, 0x2ceee22e, 0x0b54e6a6, 0xd073c03a, 0x3d3f670e, 0xdb090f3a,
	0xcb73ab2d, 0x210cc211, 0x79fc9477, 0x56db66ce, 0x7607573a, 0xc56d0340, 0x0d6f50e7, 0x0f911f2a,
	0x16f5699b, 0x63123cb0, 0x0015f81b, 0xfc22cc2b, 0x6594c4ba, 0x1d645134, 0x8633c3c5, 0x6565d5d9,
	0xc902200b, 0x8ea7aa6e, 0xa28b3d86, 0x9f22ef15, 0x9e80e834, 0x1931d611, 0xd25095ed, 0xdce57608,
	0xbe54d17a, 0xb75b7b77, 0xff53c715, 0x6d1fe6f3, 0xf4f1e1e8, 0x507749b1, 0x0c153db4, 0x7e80ad1c,
	0xa5791026, 0xad3dbe27, 0x7a65a28f, 0x9361771b, 0x570cc089, 0x8d3412aa, 0xa68fd2e0, 0xdab72770,
	0x2a303edc, 0x6477e936, 0x16f913e0, 0x09274ed9, 0xe49a321d, 0x1e64052e, 0x74ab96c9, 0xd5fdd822,
	0x3db27bd0, 0x13e13918, 0xd083f603, 0xa4cc1cd1, 0x2ff33194, 0x8f610ab0, 0xa1472c0f, 0x618f44d7,
	0x25294eab, 0x4d6915bf, 0xfce933d0, 0x32454a0a, 0xa0bdc3a7, 0xa5e7417c, 0x736be207, 0xe1859393,
	0x4b2ba3ca, 0x689c8713, 0xa1431a31, 0xb1e88845, 0xf1ab868b, 0x5a832c62, 0xb774e1ea, 0xf334763c,
	0x1692aa49, 0xdebb4312, 0x934b30b3, 0x551e3eed, 0x7e832f92, 0x73e7df4a, 0x0e51b5eb, 0xefa0c479,
	0x08804adf, 0x770ee5f0, 0x3f35314a, 0x9e2cabcc, 0x40c2f1e4, 0xe9764a79, 0xe947e751, 0x52261a4d,
	0x8c0a9ee8, 0x23e5d212, 0x954e09e5, 0xcd1af9f0, 0x23b48f97, 0x5a1a7ddc, 0xc4d467cf, 0x8a1301d3,
	0x30a40ae0, 0xdc9b40a1, 0x102bfb9f, 0x5a429b7f, 0xb0025e38, 0x58d3215e, 0xcd199bdb, 0x6738e9bd,
	0xd063b1f4, 0xf72ffc51, 0x56c10096, 0xa7959937, 0xa9e12b93, 0x40c42ab1, 0xa812d5ca, 0x712a414e,
	0x55242b16, 0x3c1e0ad7, 0x069b7f70, 0xf7b3e6c8, 0x5a592aa1, 0x84438ca2, 0xbc775fd6, 0xa9b80bd7,
	0x089bad81, 0x0d8de9cc, 0xc8b58cc9, 0xb35975c1, 0x5b39b997, 0xbff2c526, 0xb4256eb5, 0x71675891,
	0x6fbe1984, 0x306519f6, 0x08ce4519, 0xf2357abe, 0x3fc05c11, 0x30c6e91d, 0x7763fda3, 0xfdd5d266,
	0x110b6f90, 0x1f2efd86, 0x98d90a21, 0xae8eddec, 0xa2e88e17, 0xdf6d25d9, 0xb783c519, 0xff880b82,
	0x3bf0c612, 0x2bd6849c, 0x7354b07a, 0x020b7961, 0xeba8e89e, 0x2ed7d4bf, 0x8f438e34, 0xf14b33e9,
	0xe6fe502f, 0xbf986a6e, 0xa103993a, 0x27c5b0ff, 0x3abb8ca0, 0x86edf8d4, 0xd01e172e, 0x38f4a865,
	0x0dae791a, 0x1c89748f, 0xeb3e3795, 0xbfe7d73b, 0x4ec6c12a, 0x877ef600, 0x5a3cbc36, 0x116030c8,
	0xd5b7a87c, 0x524d84d9, 0x23e3e04f, 0x78097fa7, 0xfec92e57, 0x7e4db0c5, 0x3b66d2c0, 0x2ddef511,
	0x3ed80c4b, 0x13a4087f, 0x0d5ee881, 0xad6ad02e, 0x5a542426, 0x2bdef8e7, 0x446a7da7, 0xfc268a55,
	0x5d9d00bd, 0x3710d1b5, 0x270f7612, 0x38f22c86, 0xffbfec26, 0x9482aa51, 0x8dd6673b, 0x8f7c80ef,
	0x5c12531f, 0x86ae5611, 0x9cccd007, 0x4d29cbf6, 0x8a0ff3a8, 0xf0f2332d, 0x275d7034, 0xda8f94fd,
	0x5ac736fa, 0xb4cb60e4, 0x1e74c5a9, 0x53cc5ac5, 0xec538437, 0x825489d9, 0x0ba43378, 0x07657513,
	0x35ec8375, 0x1da2a732, 0x7a3b5ede, 0xab6fd84e, 0x6f8b7eda, 0x39994295, 0xd45f7faf, 0xbf6ae7c4,
	0xe4257c3d, 0x5ee315a1, 0x0bb321c5, 0x0e88401b, 0xb7053e8b, 0xd25e9808, 0x9ff33ef5, 0x89a0bd64,
	0xffdb0f83, 0xa34404c9, 0x70c36e1e, 0x9be9babb, 0x2a932500, 0x5750fd0e, 0xa4cab6f5, 0x9ec00d66,
	0x1b5f057d, 0xc88a5a6b, 0x57e3d177, 0xbc09b7d8, 0xb7eba4d3, 0x077f3fe7, 0xf8dc24f4, 0x25e5cf54,
	0xd052aef5, 0x30c74026, 0xfd5e2773, 0xce327753, 0xcabd0692, 0xcf4c8be0, 0x3af2851f, 0xf2b8cc7c,
	0x2838c54b, 0xbd2729db,
}

var s3 = [256]uint32{
	0xc570a03c, 0x1cd9298d, 0x53ac5593, 0x5cb35e31, 0xca7f4500, 0x868e31f8,
	0x68bf5639, 0x927bb899, 0x97869f8c, 0x22c8aff2, 0xe97ab5ac, 0xc4e199f7, 0x11f56e63, 0x316e6f9b,
	0xdc0b25b0, 0x3c0e37bf, 0x2260ab3d, 0xc7f5e4fe, 0x3d408195, 0x618dc6c1, 0x8801c70e, 0xc181139d,
	0xeecfb730, 0x19f23de1, 0xd9c4ed07, 0x6e4c91a3, 0x4b7131fd, 0x882fd1b0, 0x95dac0a1, 0xc764f41b,
	0xe8b192a7, 0x8c8ab9c3, 0x035446cb, 0xc8655163, 0xf6ca7757, 0xfa554923, 0x850adb81, 0x9f44293c,
	0x06742262, 0x872a79d3, 0xcc79e9fd, 0xbdaf5759, 0xa75653bb, 0x25a1c64f, 0x33bf5313, 0xfcc408f4,
	0xc61dbe73, 0x2da095d3, 0xda93f942, 0x2807d44a, 0x6663b694, 0x1383c9a1, 0xfccf0b6c, 0x2ebe6ec6,
	0x3edf77ec, 0x066d5d70, 0x6bc67bb5, 0xc54ee732, 0xebe6e605, 0xa5f5cf9a, 0x3d6c0ab2, 0x572be8c0,
	0xf02195c5, 0xcc75ff05, 0x5454bce7, 0xed431c7a, 0x35ff8d73, 0xa69f1357, 0xbe3322df, 0x8d5701d3,
	0x8227c6e1, 0x7f92b847, 0x17503b18, 0xfebf088f, 0xb969378c, 0x80695378, 0x6eb6c428, 0xf6ae7809,
	0xf8115237, 0x72659f3d, 0x90dd9052, 0xf60e6b5b, 0xe98a45d4, 0x6ca89b02, 0x85327733, 0x7c899229,
	0x923fceda, 0x987066d2, 0x3497e625, 0x3e04c58d, 0xb1be8db6, 0x172baef9, 0x30c3cc5a, 0x573dde84,
	0x67f06558, 0x8e21ff58, 0x00f3f92d, 0x6cf4cfc2, 0x13415015, 0xf461cd1d, 0xad6c6355, 0x92bf842c,
	0x274e705e, 0xee44fd1d, 0x05fd79b4, 0x40741777, 0x70a40bf2, 0x261632c0, 0xd3dae96b, 0xe8ebc9bd,
	0x3be3d490, 0xb4530a30, 0xba6dfde5, 0x3a648e2d, 0xb14c4f26, 0x7c7d0a3e, 0x559fb601, 0x44b1a722,
	0x72fcff7f, 0xf62f6adf, 0xaec6f92e, 0x6511ad20, 0x4af6ad4a, 0xaa5b3a09, 0x5303b2be, 0xbb66df75,
	0xa2490b13, 0xeacf61ba, 0x73b29c61, 0x509a66ee, 0x8080bdda, 0x9216daca, 0xfaefc031, 0x65896009,
	0x3fa36cfc, 0x995feff2, 0xce98eab5, 0x66d7e0cd, 0xe5a71216, 0xd182bc77, 0xc7d769a6, 0xda5ecc66,
	0x0473072c, 0xe84b6cc7, 0x8bbd0177, 0x0d1075aa, 0x2bf0168c, 0xa7229229, 0xbb80827b, 0xf0066c50,
	0x5a614bf6, 0x23afe56a, 0x067daa78, 0xbf01eee6, 0x5b081768, 0x1cc2f422, 0xfb6a0382, 0xa5a777a7,
	0x7609e111, 0x77097c89, 0x075c4fbf, 0x51e9004f, 0xec84f0cb, 0x1de8cc73, 0x2a54a800, 0x09a89025,
	0xfa5f8045, 0xc29b195d, 0xcff9bfe6, 0x522dbff5, 0x9c374800, 0x347dcd8f, 0x974da9c0, 0x8d6a6d88,
	0xb47ef442, 0xa51e66ca, 0xa210c54a, 0x4f63c725, 0xff1a465d, 0x813eb31b, 0x0e0058d5, 0x3c18ce5a,
	0xc4d7d98c, 0x4e24da16, 0x5aec5af6, 0x912cd19f, 0xb12bf2d1, 0x184d3b0b, 0x82dbd6da, 0xd29ead22,
	0x9d13dce5, 0xbea27f78, 0xb957b027, 0xe0dae424, 0x1ae3ab8f, 0x49c349a2, 0x74efda3d, 0x88539bd4,
	0xb9f027c3, 0x5739e997, 0x08d6028e, 0x8d1f0b8f, 0x63256408, 0x9b216118, 0xd89432d3, 0x3bebf6ca,
	0x21735953, 0x0eda4bfb, 0xe6afc2d4, 0xa9db95f9, 0x1f1c6bb0, 0xbaaf121b, 0x8cdc1b36, 0x3913f9fd,
	0x863bcb1a, 0xbd34adcb, 0xda48457a, 0x4f584129, 0xdd85156c, 0x0324f396, 0xd41e1ee1, 0xc3b48f82,
	0x2124fb4c, 0x6c0b2635, 0x95ce3157, 0xa8daca8c, 0xb54e1542, 0xd989f76a, 0x0c1ea5e3, 0x973fe85c,
	0xe6e91d97, 0x2916b8df, 0x5b0b05e2, 0x57bdc906, 0x7cb2ccef, 0x131c7553, 0x41ca9311, 0x6e70e1c1,
	0x0f972bf2, 0x8cf59d7a, 0xe6613022, 0x69218e19, 0xc350744a, 0xa2d5bf1b, 0x9fa14b7b, 0x8867f25c,
	0xb8e19af6, 0x777a25df, 0xa2004e28, 0xfb929664, 0x2da8a284, 0x21955d47, 0xf0cebd06, 0x9887b1e9,
	0xbf7810c1, 0x265d91f9,
}

var p = [18]uint32{
	0x640cded2, 0xca6cf7cf, 0xc7bc95fb, 0x7d0d60a3, 0xcf23ad88, 0x8ffb62dc, 0x6c3da5cc, 0x6bfcd6d6,
	0x63f492df, 0xe32ebe65, 0xc3746b6d, 0xc5703934, 0xdc940bce, 0x590e0892, 0xea9413e8, 0xf4b13de7,
	0x505893fc, 0xe3d696e3,
}
