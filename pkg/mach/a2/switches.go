package a2

// Here we set up all the soft switches that we'll use in the computer,
// which is a lot!
func (c *Computer) defineSoftSwitches() {
	c.MapRange(0x0, 0x200, zeroPageRead, zeroPageWrite)
	c.MapRange(0x0400, 0x0800, displayRead, displayWrite)
	c.MapRange(0x2000, 0x4000, displayRead, displayWrite)
	c.MapRange(0xC100, 0xD000, pcRead, pcWrite)
	c.MapRange(0xD000, 0x10000, bankRead, bankWrite)

	msc := newMemorySwitchCheck()
	c.RMap[0xC013] = msc.IsSetter(MemReadAux)
	c.RMap[0xC014] = msc.IsSetter(MemWriteAux)
	c.RMap[0xC018] = msc.IsSetter(Mem80Store)
	c.RMap[0xC01C] = msc.IsSetter(MemPage2)
	c.RMap[0xC01D] = msc.IsSetter(MemHires)
	c.RMap[0xC054] = msc.UnSetterR(MemPage2)
	c.RMap[0xC055] = msc.ReSetterR(MemPage2)
	c.RMap[0xC056] = msc.UnSetterR(MemHires)
	c.RMap[0xC057] = msc.ReSetterR(MemHires)
	c.WMap[0xC000] = msc.UnSetterW(Mem80Store)
	c.WMap[0xC001] = msc.ReSetterW(Mem80Store)
	c.WMap[0xC002] = msc.UnSetterW(MemReadAux)
	c.WMap[0xC003] = msc.ReSetterW(MemReadAux)
	c.WMap[0xC004] = msc.UnSetterW(MemWriteAux)
	c.WMap[0xC005] = msc.ReSetterW(MemWriteAux)
	c.WMap[0xC054] = msc.UnSetterW(MemPage2)
	c.WMap[0xC055] = msc.ReSetterW(MemPage2)
	c.WMap[0xC056] = msc.UnSetterW(MemHires)
	c.WMap[0xC057] = msc.ReSetterW(MemHires)

	bsc := newBankSwitchCheck()
	c.RMap[0xC080] = bsc.SetterR(BankRAM | BankRAM2)
	c.RMap[0xC081] = bsc.SetterR(BankWrite | BankRAM2)
	c.RMap[0xC082] = bsc.SetterR(BankRAM2)
	c.RMap[0xC083] = bsc.SetterR(BankRAM | BankWrite | BankRAM2)
	c.RMap[0xC088] = bsc.SetterR(BankRAM)
	c.RMap[0xC089] = bsc.SetterR(BankWrite)
	c.RMap[0xC08A] = bsc.SetterR(BankDefault)
	c.RMap[0xC08B] = bsc.SetterR(BankRAM | BankWrite)
	c.RMap[0xC011] = bsc.IsSetter(BankRAM2)
	c.RMap[0xC012] = bsc.IsSetter(BankRAM)
	c.RMap[0xC016] = bsc.IsSetter(BankAuxiliary)
	c.WMap[0xC008] = bsc.UnSetterW(BankAuxiliary)
	c.WMap[0xC009] = bsc.ReSetterW(BankAuxiliary)

	psc := newPCSwitchCheck()
	c.RMap[0xC015] = psc.IsSetter(PCSlotCxROM)
	c.RMap[0xC017] = psc.IsSetter(PCSlotC3ROM)
	c.WMap[0xC006] = psc.ReSetterW(PCSlotCxROM)
	c.WMap[0xC007] = psc.UnSetterW(PCSlotCxROM)
	c.WMap[0xC00A] = psc.UnSetterW(PCSlotC3ROM)
	c.WMap[0xC00B] = psc.ReSetterW(PCSlotC3ROM)

	dsc := newDisplaySwitchCheck()
	c.RMap[0xC01A] = dsc.IsSetter(DisplayText)
	c.RMap[0xC01B] = dsc.IsSetter(DisplayMixed)
	c.RMap[0xC01E] = dsc.IsSetter(DisplayAltCharset)
	c.RMap[0xC01F] = dsc.IsSetter(Display80Col)
	c.RMap[0xC050] = dsc.UnSetterR(DisplayText)
	c.RMap[0xC051] = dsc.ReSetterR(DisplayText)
	c.RMap[0xC052] = dsc.UnSetterR(DisplayMixed)
	c.RMap[0xC053] = dsc.ReSetterR(DisplayMixed)
	// Technically DHires should only be reset/unset if IOU is set; I'm
	// not sure if this is necessary in practice.
	c.RMap[0xC05E] = dsc.ReSetterR(DisplayDHires)
	c.RMap[0xC05F] = dsc.UnSetterR(DisplayDHires)
	// --
	c.RMap[0xC07E] = dsc.IsSetter(DisplayIOU)
	c.RMap[0xC07F] = dsc.IsSetter(DisplayDHires)
	c.WMap[0xC00C] = dsc.UnSetterW(Display80Col)
	c.WMap[0xC00D] = dsc.ReSetterW(Display80Col)
	c.WMap[0xC00E] = dsc.UnSetterW(DisplayAltCharset)
	c.WMap[0xC00F] = dsc.ReSetterW(DisplayAltCharset)
	c.WMap[0xC050] = dsc.UnSetterW(DisplayText)
	c.WMap[0xC051] = dsc.ReSetterW(DisplayText)
	c.WMap[0xC052] = dsc.UnSetterW(DisplayMixed)
	c.WMap[0xC053] = dsc.ReSetterW(DisplayMixed)
	c.WMap[0xC05E] = dsc.ReSetterW(DisplayDHires)
	c.WMap[0xC05F] = dsc.UnSetterW(DisplayDHires)
	c.WMap[0xC07E] = dsc.ReSetterW(DisplayIOU)
	c.WMap[0xC07F] = dsc.UnSetterW(DisplayIOU)
}
