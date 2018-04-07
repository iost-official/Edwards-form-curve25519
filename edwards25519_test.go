package edwards25519

// This test code is used to generate testvectors for  Rust curve25519-dalek.

import "testing"
import "fmt"

// The additive identity added to it should produce itself.
func Test_ProjectiveGroupElement_zero(t *testing.T) {
	var p1 ProjectiveGroupElement;
	var b  [32]byte;
	var p2 ExtendedGroupElement;
	var p3 CompletedGroupElement;
	var p4 ProjectiveGroupElement;
	var p5 ExtendedGroupElement;
	var p6 CachedGroupElement;

	p1.Zero()
	p1.ToBytes(&b)
	p5.FromBytes(&b)
	p5.ToCached(&p6)

	p2.Zero()
	geAdd(&p3, &p2, &p6)

	p3.ToProjective(&p4)

	fmt.Println("Test_ProjectiveGroupElement_zero")
	fmt.Println(p1)
	fmt.Println(p3)
	fmt.Println(p4)
}

func Test_ProjectiveGroupElement_double_identity(t *testing.T) {
	var p1 ProjectiveGroupElement;
	var p2 CompletedGroupElement;

	p1.Zero()
	p1.Double(&p2)

	fmt.Println("Test_ProjectiveGroupElement_double_identity")
	fmt.Println(p2)
}

func Test_ProjectiveGroupElement_double(t *testing.T) {
	var t1 ExtendedGroupElement;
	var t2 PreComputedGroupElement;

	var p1 CompletedGroupElement;
	var p2 ProjectiveGroupElement;
	var p3 CompletedGroupElement;

	t1.Zero()
	selectPoint(&t2, 1, -1)

	geMixedAdd(&p1, &t1, &t2)
	p1.ToProjective(&p2)
	p2.Double(&p3)

	fmt.Println("Test_ProjectiveGroupElement_double")
	fmt.Println(p3)
}

func Test_ExtendedGroupElement_add_cached(t *testing.T) {
	var t1 ExtendedGroupElement;
	var t2 PreComputedGroupElement;

	var p1 CompletedGroupElement;
	var p2 ExtendedGroupElement;
	var p3 CachedGroupElement;
	var p4 CompletedGroupElement;

	t1.Zero()
	selectPoint(&t2, 1, -1)  // p1

	geMixedAdd(&p1, &t1, &t2)
	p1.ToExtended(&p2)
	p2.ToCached(&p3);
	geAdd(&p4, &p2, &p3)

	fmt.Println("Test_ExtendedGroupElement_add_cached")
	fmt.Println(p3)
	fmt.Println(p4)
}

func Test_ExtendedGroupElement_sub_cached(t *testing.T) {
	var t1 ExtendedGroupElement;
	var t2 PreComputedGroupElement;

	var p1 CompletedGroupElement;
	var p2 ExtendedGroupElement;
	var p3 CachedGroupElement;
	var p4 CompletedGroupElement;

	t1.Zero()
	selectPoint(&t2, 1, -1)  // p1

	geMixedAdd(&p1, &t1, &t2)
	p1.ToExtended(&p2)
	p2.ToCached(&p3);
	geSub(&p4, &p2, &p3)

	fmt.Println("Test_ExtendedGroupElement_sub_cached")
	fmt.Println(p3)
	fmt.Println(p4)
}

func Test_ExtendedGroupElement_add_precomputed(t *testing.T) {
	var t1 ExtendedGroupElement;
	var t2 PreComputedGroupElement;

	var p1 CompletedGroupElement;
	var p2 ExtendedGroupElement;
	var p3 PreComputedGroupElement;
	var p4 CompletedGroupElement;

	t1.Zero()
	selectPoint(&t2, 1, -1)  // p1

	geMixedAdd(&p1, &t1, &t2)
	p1.ToExtended(&p2)
	selectPoint(&p3, 27, -3);
	geMixedAdd(&p4, &p2, &p3)

	fmt.Println("Test_ExtendedGroupElement_add_precomputed")
	fmt.Println(p3)
	fmt.Println(p4)
}

func Test_ExtendedGroupElement_sub_precomputed(t *testing.T) {
	var t1 ExtendedGroupElement;
	var t2 PreComputedGroupElement;

	var p1 CompletedGroupElement;
	var p2 ExtendedGroupElement;
	var p3 PreComputedGroupElement;
	var p4 CompletedGroupElement;

	t1.Zero()
	selectPoint(&t2, 1, -1)  // p1

	geMixedAdd(&p1, &t1, &t2)
	p1.ToExtended(&p2)
	selectPoint(&p3, 2, -5);
	geMixedSub(&p4, &p2, &p3)

	fmt.Println("Test_ExtendedGroupElement_sub_precomputed")
	fmt.Println(p3)
	fmt.Println(p4)
}

func Test_selectPoint_27_negative3(t *testing.T) {
	var p1 PreComputedGroupElement;

	selectPoint(&p1, 27, -3)

	fmt.Println("Test_selectPoint_27_negative3")
	fmt.Println(p1)
}

func Test_selectPoint_29_negative8(t *testing.T) {
	var p1 PreComputedGroupElement;

	selectPoint(&p1, 29, -8)

	fmt.Println("Test_selectPoint_29_negative8")
	fmt.Println(p1)
}