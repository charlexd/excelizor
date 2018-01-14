local SimpleTypes = {
    [1001] = {
        Id = 1001,
        NumberTest = 1,
        StringTest = "这是测试字符串1",
        FloatTest = 8.900,
        BoolTest = true, 
        DictTest1 = {
            Item1 = 3.330,
            Item2 = 4.440,
            Item3 = 5.550, 
        }, 
        DictTest2 = {
            Item1 = 2.220,
            Item2 = 3.330, 
        },
        ListTest1 = {
            [0] = "第一个字符串",
            [1] = "第二个字符串",
            [2] = "第一个字符串",
        },
        ListTest2 = {
            [0] = "第一个字符串",
            [1] = "second string",
            [2] = "third string",
            [3] = "",
        }, 
        DictTest3 = { 
            SubDict1 = {
                Item1 = "DictTest3[\"SubDIct1\"][\"Item1\"}",
                Item2 = "DictTest3[\"SubDIct1\"][\"Item2\"}", 
            }, 
            SubDict2 = {
                Item1 = "DictTest3[\"SubDIct2\"][\"Item1\"}",
                Item2 = "DictTest3[\"SubDIct2\"][\"Item2\"}", 
            }, 
        }, 
        DictTest4 = { 
            SubDict1 = {
                item1 = "seges",
                item2 = "fesgs",
                item3 = "fesge", 
            }, 
            SubDict2 = {
                item1 = "123",
                item2 = "121", 
            }, 
        }, 
        DictTest5 = { 
            Subdict1 = {
                item1 = "asd",
                item2 = "sdf", 
            }, 
            Subdict2 = {
                item1 = "qwe",
                item2 = "wer", 
            }, 
        },
        ListTest3 = {
            [0] = {
                [0] = "ListTest3[0][0]",
                [1] = "LIstTest3[0][1]",
            },
            [1] = {
                [0] = "ListTest3[1][0]",
                [1] = "ListTest3{1][1]",
            },
        },
        ListTest4 = {
            [0] = {
                [0] = "list[0][0]",
                [1] = "List[0][1]",
            },
            [1] = {
                [0] = "list[1][0]",
                [1] = "list[1][1]",
            },
        },
        ListTest5 = {
            [0] = {
                [0] = "123",
                [1] = "234",
                [2] = "345",
            },
            [1] = {
                [0] = "234",
                [1] = "345",
                [2] = "456",
            },
        }, 
        DictTest6 = {
            sublist1 = {
                [0] = "dictTest6[\"sublist1\"][0]",
                [1] = "dictTest6[\"sublist1\"][1]",
            },
            sublist2 = {
                [0] = "dictTest6[\"sublist2\"][0]",
                [1] = "dictTest6[\"sublist2\"][1]",
            }, 
        }, 
        DictTest7 = {
            sublist1 = {
                [0] = "zxc",
                [1] = "xcv",
                [2] = "cvb",
            },
            sublist2 = {
                [0] = "vbn",
                [1] = "bnm",
                [2] = "mnb",
            }, 
        },
        ListTest6 = { 
            [0] = {
                item1 = "fse",
                item2 = "fsd", 
            }, 
            [1] = {
                item1 = "tyu",
                item2 = "poi", 
            },
        },
        ListTest7 = { 
            [0] = {
                item0_1 = "fse",
                item0_2 = "fsd", 
            }, 
            [1] = {
                item1_1 = "qwe",
                item1_2 = "qse", 
            },
        },
    },
    [1002] = {
        Id = 1002,
        NumberTest = 2,
        StringTest = "这是测试字符串2",
        FloatTest = 9.100,
        BoolTest = false, 
        DictTest1 = {
            Item1 = 4.440,
            Item2 = 5.550,
            Item3 = 6.660, 
        }, 
        DictTest2 = {
            Item1 = 2.220,
            Item2 = 3.330, 
        },
        ListTest1 = {
            [0] = "第一个字符串",
            [1] = "第二个字符串",
            [2] = "第一个字符串",
        },
        ListTest2 = {
            [0] = "第一个字符串",
            [1] = "second string",
            [2] = "third string",
            [3] = "",
        }, 
        DictTest3 = { 
            SubDict1 = {
                Item1 = "DictTest3[\"SubDIct1\"][\"Item1\"}",
                Item2 = "DictTest3[\"SubDIct1\"][\"Item2\"}", 
            }, 
            SubDict2 = {
                Item1 = "DictTest3[\"SubDIct2\"][\"Item1\"}",
                Item2 = "DictTest3[\"SubDIct2\"][\"Item2\"}", 
            }, 
        }, 
        DictTest4 = { 
            SubDict1 = {
                item1 = "seges",
                item2 = "fesgs",
                item3 = "fesge", 
            }, 
            SubDict2 = {
                item1 = "123",
                item2 = "121", 
            }, 
        }, 
        DictTest5 = { 
            Subdict1 = {
                item1 = "asd",
                item2 = "sdf", 
            }, 
            Subdict2 = {
                item1 = "qwe",
                item2 = "wer", 
            }, 
        },
        ListTest3 = {
            [0] = {
                [0] = "ListTest3[0][0]",
                [1] = "LIstTest3[0][1]",
            },
            [1] = {
                [0] = "ListTest3[1][0]",
                [1] = "ListTest3{1][1]",
            },
        },
        ListTest4 = {
            [0] = {
                [0] = "list[0][0]",
                [1] = "List[0][1]",
            },
            [1] = {
                [0] = "list[1][0]",
                [1] = "list[1][1]",
            },
        },
        ListTest5 = {
            [0] = {
                [0] = "123",
                [1] = "234",
                [2] = "345",
            },
            [1] = {
                [0] = "234",
                [1] = "345",
                [2] = "456",
            },
        }, 
        DictTest6 = {
            sublist1 = {
                [0] = "dictTest6[\"sublist1\"][0]",
                [1] = "dictTest6[\"sublist1\"][1]",
            },
            sublist2 = {
                [0] = "dictTest6[\"sublist2\"][0]",
                [1] = "dictTest6[\"sublist2\"][1]",
            }, 
        }, 
        DictTest7 = {
            sublist1 = {
                [0] = "zxc",
                [1] = "xcv",
                [2] = "cvb",
            },
            sublist2 = {
                [0] = "vbn",
                [1] = "bnm",
                [2] = "mnb",
            }, 
        },
        ListTest6 = { 
            [0] = {
                item1 = "fse",
                item2 = "fsd", 
            }, 
            [1] = {
                item1 = "tyu",
                item2 = "poi", 
            },
        },
        ListTest7 = { 
            [0] = {
                item0_1 = "fse",
                item0_2 = "fsd", 
            }, 
            [1] = {
                item1_1 = "qwe",
                item1_2 = "qse", 
            },
        },
    },
    [1003] = {
        Id = 1003,
        NumberTest = 2,
        StringTest = "这是测试字符串2",
        BoolTest = false, 
        DictTest1 = {
            Item1 = 4.440,
            Item3 = 6.660, 
        }, 
        DictTest2 = {
            Item1 = 2.220,
            Item2 = 3.330, 
        },
        ListTest1 = {
            [0] = "第一个字符串",
            [1] = "第一个字符串",
        },
        ListTest2 = {
            [0] = "第一个字符串",
            [1] = "second string",
            [2] = "third string",
            [3] = "",
        }, 
        DictTest3 = { 
            SubDict1 = {
                Item1 = "DictTest3[\"SubDIct1\"][\"Item1\"}",
                Item2 = "DictTest3[\"SubDIct1\"][\"Item2\"}", 
            }, 
            SubDict2 = {
                Item1 = "DictTest3[\"SubDIct2\"][\"Item1\"}", 
            }, 
        }, 
        DictTest4 = { 
            SubDict1 = {
                item1 = "seges",
                item2 = "fesgs",
                item3 = "fesge", 
            }, 
            SubDict2 = {
                item1 = "123",
                item2 = "121", 
            }, 
        }, 
        DictTest5 = { 
            Subdict1 = {
                item1 = "asd",
                item2 = "sdf", 
            }, 
            Subdict2 = {
                item1 = "qwe",
                item2 = "wer", 
            }, 
        },
        ListTest3 = {
            [0] = {
                [0] = "ListTest3[0][0]",
                [1] = "LIstTest3[0][1]",
            },
            [1] = {
                [0] = "ListTest3[1][0]",
                [1] = "ListTest3{1][1]",
            },
        },
        ListTest4 = {
            [0] = {
                [0] = "list[0][0]",
                [1] = "List[0][1]",
            },
            [1] = {
                [0] = "list[1][0]",
                [1] = "list[1][1]",
            },
        },
        ListTest5 = {
            [0] = {
                [0] = "123",
                [1] = "234",
                [2] = "345",
            },
            [1] = {
                [0] = "234",
                [1] = "345",
                [2] = "456",
            },
        }, 
        DictTest6 = {
            sublist1 = {
                [0] = "dictTest6[\"sublist1\"][0]",
                [1] = "dictTest6[\"sublist1\"][1]",
            },
            sublist2 = {
                [0] = "dictTest6[\"sublist2\"][0]",
                [1] = "dictTest6[\"sublist2\"][1]",
            }, 
        }, 
        DictTest7 = {
            sublist1 = {
                [0] = "zxc",
                [1] = "xcv",
                [2] = "cvb",
            },
            sublist2 = {
                [0] = "vbn",
                [1] = "bnm",
                [2] = "mnb",
            }, 
        },
        ListTest6 = { 
            [0] = {
                item1 = "fse",
                item2 = "fsd", 
            }, 
            [1] = {
                item1 = "tyu",
                item2 = "poi", 
            },
        },
        ListTest7 = { 
            [0] = {
                item0_1 = "fse",
                item0_2 = "fsd", 
            }, 
            [1] = {
                item1_1 = "qwe",
                item1_2 = "qse", 
            },
        },
    },
}

return SimpleTypes