package camelcard

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	bid   int
	cards string
	power int
}

var INPUT = "AATKJ 840\n27A83 251\n6TT8T 113\nQ6T6T 693\n99K89 553\n777Q7 136\n7227J 782\nTJ2J9 837\n55578 524\n24Q49 919\nAAA2A 709\nKT972 907\n85282 886\n4432J 255\nK6K27 119\nJ9Q77 820\n33K3K 48\nT8887 500\n28272 77\n38Q46 287\nKQQAQ 481\n5T2T3 280\nKQQQQ 458\nK82Q3 712\nQ5552 758\n4K87K 174\n7QQJQ 331\n29233 96\n79KA3 65\nJ9893 546\nQ5A39 842\nK4K7K 437\nT555T 215\n9A7AA 162\n32AA2 374\n8J222 134\n38376 583\nJ5222 962\nJ65A9 883\n36338 956\nQ654T 107\nK4KKK 169\nA9A76 710\n9KQ78 183\n88753 577\nK997J 539\n4QA88 914\nK5T3A 401\n8K77K 619\n8888J 400\nTATAT 58\n8T436 279\n3Q28A 659\n3KK4J 54\nQQJ59 903\n44777 779\n53336 477\n22J33 930\nTQ6T2 923\n475TK 633\n5KJ45 980\n47TK6 50\nK8KKK 453\n2TJKA 39\nTJ6K6 380\n4J888 311\n27277 317\n2A2A5 213\nTT7T3 365\n22555 858\n789QQ 421\n937KT 791\n89KT6 732\n3KKKJ 703\nA2T22 896\n6467J 489\n72878 863\n46J5J 283\nTATA6 71\nQ4967 941\n7JT55 232\n9TQQQ 785\n55445 835\nK68KK 256\nTT33T 72\n2A2AQ 435\n9643K 436\nJ3282 95\n88TQQ 767\nQ8788 250\nQQ949 62\n4T552 448\nJJ8QQ 476\n2JT22 900\n73T7T 648\nQQQJJ 795\n98A35 589\n94444 479\n44486 465\n7TK86 521\n2JK45 761\n556KT 492\n94K99 144\n2QJ32 637\n66226 538\n933TA 248\nT2234 419\nQAA26 35\nK63K3 853\n5KJK9 970\nKK4QK 636\n8TTTT 69\n5T7J8 161\n239A4 91\n3T823 182\n24A8K 407\nK5A92 330\nKK9K3 599\n2244J 284\nAAJ5Q 794\n44A84 472\nKK2J3 597\n6QQJ6 547\nTAAT4 668\nJ4444 827\nQAJ22 200\nKT5J5 80\n7K3J9 517\n3A33T 762\nTT66T 154\n6K972 569\n6378Q 593\n2AQA9 460\n2KKAA 516\nAK45A 978\nQAAA8 568\n83478 234\n33737 491\nQ9Q43 977\nQKQQ7 523\n9T24A 609\n88997 355\nKAKAA 651\n87Q77 665\nJK7A4 268\n68586 44\n4345J 931\nQ29A9 871\n3479K 774\n3TK65 189\n9249Q 319\n27J2Q 656\n333J9 667\n6TJ66 681\n82T96 738\nK3K39 14\n74JK9 297\n8A88A 798\nAT869 987\n9TAT9 996\n33939 684\n4424J 117\n44AJA 422\n646AA 358\nK3T83 218\nKK7K7 495\nA6A6J 731\n3QQ97 381\n43ATK 484\nK9965 982\n9J9JJ 392\nQ2QQQ 420\n58588 433\n9Q955 862\n88777 246\n669J7 438\n9AA2A 733\n8A9K6 175\nQ25JK 459\nJJ2K8 439\nT8K9T 195\n45444 719\nA22A2 403\n6366Q 36\n35366 452\n293J3 576\nJKK3J 402\n425AT 338\n344J4 975\n7K7A8 490\n7QQ7Q 242\nT9T9K 542\n2JJT6 440\n76676 541\nAKA66 361\n7557Q 431\n5K552 511\n997K7 179\n79779 885\n85585 961\nTQQQ8 751\n9TAT6 504\n84248 654\n222JJ 865\n6AKK8 573\nAQ3T8 30\n77676 670\nT44J2 808\nJ7J5Q 897\n44343 57\nTJ7T5 945\nJKTTK 60\n3333Q 894\n6886J 473\n8J7J3 990\n77797 915\n77J98 155\nKKQT8 736\n8A8AJ 219\nAQQJA 75\n75J58 349\n78978 898\n48TAA 726\n56A6A 45\n888JJ 170\nQ7KK3 294\nQQA7A 199\nA4653 193\n44545 602\n622K7 612\nT5TJ9 860\n8T888 548\n8TQJ4 410\nTQ578 446\n98988 367\n5989T 151\nQQ35J 679\n73677 416\nJ33Q5 137\n6564J 913\n4847Q 807\n39292 874\nJQA45 810\n97975 943\nQ9839 296\n9599T 661\nJ9699 723\n44J72 875\nT29A9 889\nQQ53K 1000\n88988 20\nQJQK5 769\nKAKKK 132\n7327J 177\nQQQAQ 704\n2KTK6 257\n474A4 928\n9A748 501\n5A535 469\n2T426 190\n53353 318\nQ7Q78 694\n92738 796\n25995 772\nQQQ22 877\n9K999 238\nJJQ5A 424\n4QA4J 7\nKQT2A 347\n27787 754\n823T5 159\n54347 702\nATA55 166\n9Q8T7 887\nK8338 631\n945AJ 658\nTJTTJ 16\n7A478 313\n93998 2\nA5A93 342\nKKQQK 921\n3K242 52\n9AAJT 43\nTT33Q 337\n22452 139\n2AQK8 594\n6AQ58 503\nKJ5JQ 122\n9Q787 988\n7626T 228\nJ99J9 377\n6J727 378\n45J33 216\nA555K 360\n6J62K 959\nQ7Q47 141\n95644 116\n5K755 56\n7A7QQ 773\n27K92 663\nA4A77 841\n763J4 872\n3493J 478\n42K9K 103\n43433 973\nK6A4K 38\nT257Q 307\n2KQKK 992\nTT553 627\n66664 508\n88A82 243\nQJQQQ 804\n2T99T 150\n2JQT3 245\n44644 786\n675TK 194\n45KKA 1\nJJJJJ 543\nT4TTT 530\nK6KK6 336\nTT8Q3 888\nT999Q 350\n36643 176\n95KKK 832\nJ7386 375\nK8KK8 954\n82JK3 411\nAAAAT 571\n3677J 854\n9998J 604\n5JQT4 94\nJAKKQ 635\nTT986 15\n5855Q 756\n243Q8 127\nT69A3 821\n63A7T 902\nQQ8Q8 805\nAQAAK 371\n83J33 981\nAQ525 549\n72486 147\n36664 715\nQQTTQ 312\n686A9 920\nJ63J7 486\n4AAAQ 451\nQ5452 84\n66688 191\n3QAQ9 285\n925K6 66\n824Q8 5\nJTQQT 806\nJ5J33 344\nA2A35 428\n72943 262\nK9966 581\n66J69 991\n788J7 728\nAA7AA 471\n5555K 167\nJ4A62 488\nJJ699 258\n36386 600\nA4AA4 405\n53Q9T 211\n3377A 985\nK7Q83 404\nJ3TT7 265\nQ9QQQ 482\n82736 273\n92K4A 771\n68888 417\nA7AA7 625\n29333 630\n558QK 536\n7J9J7 247\n66T66 586\nAKA9A 11\n73J3J 67\nQ6226 570\n885J5 457\n7447J 869\nTA2A2 688\nJ9685 254\n4JKK4 275\n8QQQ7 442\n82383 244\n43333 19\n573JT 701\nA666A 777\n999J9 984\nJJ444 822\nKJ3A4 505\n22494 387\nTJT77 301\nK2K2K 112\n88338 844\n5AA5A 388\n33J63 233\n9658Q 432\n797K7 291\nAJ3TT 935\nK2255 483\n66966 286\n33555 101\nAQ4QQ 944\n77K77 575\n4T4QQ 650\n32J6T 948\nAK32Q 666\n57455 685\n9823A 708\n8TTTK 474\n22575 102\n77577 368\n88J28 90\nA4343 610\n24826 82\nQQ8QQ 158\n99QQ9 892\n77468 99\n3T773 4\n4QQ3Q 620\n85437 212\n53655 745\nQ6664 156\n7447Q 320\n673K5 502\n4TJ7Q 917\n4TJ44 281\n53545 186\nTQ54A 957\n8J242 818\n4J75A 938\n95T8A 499\nKJ6KK 657\nT5TTT 12\nAAA4A 878\n6Q66K 788\n7TT22 450\n882QA 880\n5KT99 397\n3955T 263\nJ68Q8 939\n6T395 742\n3KQ69 345\nQ9333 425\nJ7777 695\n722Q7 130\nT7TT6 775\nTTT88 78\nAQQ22 815\nJ8837 616\n223Q3 271\n4KKK4 28\n95999 879\nATTT6 34\n4J4JA 603\nQ5972 37\n798K6 963\n29J4Q 217\n2AAA3 830\n99J9T 893\n5J2K2 235\nTTT55 937\n9587T 906\n4T3JK 968\nK88K8 734\nJT223 801\n9QJ2Q 418\nK8866 528\n887AA 487\n6733J 426\n82887 17\n54TQK 643\n44J94 639\n58K6A 470\n3TTTT 223\nT4TT4 760\n82435 942\nK494Q 114\n77792 556\n5J82J 118\n2T8AA 559\n9494A 706\nJ4666 295\n33J7K 749\n8JQ88 717\nTJTTT 672\n4AAKK 557\nQ77QJ 629\nKKJK5 535\n492JT 834\n9QK74 876\n7QQQQ 308\n622J2 946\nJ489Q 160\nA2A2A 647\n845T2 362\nQ7TT4 104\nT7KQA 513\n8T8T8 891\n25J82 6\n3333J 53\nJ4544 443\n3377Q 551\nKKTQQ 655\n22822 59\nQ9Q9Q 278\n94JTK 363\n66699 206\n7A499 882\nT3663 677\n4K44Q 21\n82A42 197\n4824Q 566\nT8TT4 624\nKQJ7Q 951\nTJT89 163\n65A5A 884\n3T63A 92\nA3A34 434\nT7J88 202\n4K9A3 969\n52333 814\nJ3977 615\n86666 697\n8A987 866\n5434K 464\n75888 974\n4444K 115\n62K32 143\n5J5TT 140\nT4QTQ 324\nK2282 314\n39993 518\n2AJ4J 83\nK96K9 952\nQ683J 851\nJJ4J8 724\n2222T 611\n7J676 595\n3T333 540\n88848 3\n2Q679 463\nJ68J6 97\nAAA5A 276\n22272 288\n3353J 171\nKK67J 207\n34QJ7 617\nTTAKK 714\nAA6JA 529\nA8299 373\n5Q222 79\n666J6 953\nKQ29Q 799\n533TT 735\nJQAQQ 901\n995T3 716\nK2AK2 531\n7776Q 68\n94Q4Q 87\nAAA88 999\nT2434 149\nTQTQ7 385\n255T5 302\n99T99 449\n798TA 496\n7T6Q8 746\nQ2QKQ 873\n75776 300\n28588 303\n5KT5A 922\nKJT64 409\nAJAAA 924\n595JA 867\n44J93 27\n2K384 690\nJKKKK 744\nK57KQ 124\nT4464 109\n29AKJ 705\n8JK8K 261\nK6T3J 947\nT5588 125\n6A8K4 316\n35555 305\n63AAJ 560\n66J67 578\nKQ77K 646\nT654J 33\n38AKK 230\nAJ99A 838\nT6TTK 252\n6666K 689\n99A33 965\n88QAJ 764\nT9Q23 899\n8A22T 829\n478Q9 366\nTT2TT 971\n34T2A 259\nAAQAJ 23\nJ4T52 776\n9999A 277\n68896 994\n55585 18\n88999 131\n77389 514\n75AQJ 787\n88Q88 51\n3J4JQ 678\n4J643 394\n6634J 757\n74496 249\n555J5 226\nT3TT4 881\n93739 157\n377K8 328\nJJ6J3 81\n8JA9K 720\nT3TT5 828\n33A33 106\nK6KKK 413\n86868 895\n6J3QQ 323\n4JQ93 192\n939J9 729\n7Q77Q 671\nQ2222 676\n8A83Q 823\nTTT3J 74\n94969 816\n5TK77 766\nQKA87 467\nJQ736 427\nJT374 870\nJ949J 456\n8AT8A 299\nK9KK8 395\n932TA 669\nJ3AAT 105\n474KK 441\n7J738 713\n742AT 793\nK387J 145\n4457A 264\n5T5J5 908\n99996 121\nAAJ4A 817\nJ4JJJ 25\n99J59 63\n5TKKA 983\nJ454K 396\nK222J 290\n99992 393\nA2JQQ 321\n58AAA 790\n39333 165\n4264T 737\n69969 327\nAQAQA 537\nKT9J8 49\n4QT86 126\nKKAJK 304\n84448 414\nQ2425 444\n6558A 675\nA9QAJ 770\n333QQ 607\n93983 73\n929T9 856\nJJAAA 352\n67TQT 382\n6579K 55\nTT7TT 544\n9AAK6 784\n65A82 855\n48KJ9 111\n388AK 272\n6T677 890\n2AA92 632\n663J3 623\nTTJQT 386\n49475 180\n4TKTT 461\n88333 722\nQ97T4 740\n7T624 315\n66K6K 86\n33398 826\n922J5 912\n93875 89\nKK3AK 447\n32T2T 383\n77KKJ 699\n9T5A7 622\n3J823 750\n8TKJ6 224\n3464T 730\n54555 138\n526J7 389\nA5794 497\n779JK 916\nT333T 545\n74253 743\n29222 227\nTKTJT 552\n9AJ89 574\n637J6 236\nKKQKK 868\n9QQJ4 683\nK42Q6 267\n9243Q 934\n8T98T 210\nT6K58 831\n29J99 29\nJQ4QT 353\n56QTQ 692\n335J9 825\n24J84 208\nA2924 859\nA8888 933\n577K2 148\n5AKKK 429\n23KK3 201\n29229 811\n44T4T 455\n9QT2J 755\n39QAT 172\nA2548 423\nAK692 270\nA44Q4 698\nA4443 510\n69527 415\n4T444 638\nAKJ36 927\nK355A 561\n32Q3Q 220\n22A22 93\nJTQTA 850\n2Q2Q2 188\n77787 135\n68T6T 813\nAAQA7 168\nT4554 967\nK6888 240\n2AA8A 642\nAA333 13\n8KTKK 70\n99K93 686\n89588 949\nK4565 372\n8T9K4 601\n33TJJ 763\n23J33 565\nKQK65 819\nK534J 721\n43385 563\nJ5AA5 379\n88222 979\n5JTTJ 515\n2QQ8A 196\n8T756 849\n38288 454\n386QT 310\nQ5TT5 936\n4QQTQ 852\nQQ446 46\nQT6JQ 926\n242Q2 605\n44999 26\n93929 10\nA23A3 925\n66366 564\n88A8Q 809\n93222 356\n45AQ6 800\nQKKAK 359\n7A645 960\nJ2KQQ 289\nKKK5K 203\n7QA44 485\nJQQ69 128\nQ2QT5 22\nK84KK 993\nJ32A4 430\n999QK 778\n68Q3K 753\n4884J 221\nJ44JJ 340\nJTT93 339\n63TT8 348\n4TTT6 725\n6KA2Q 335\nQTKKT 592\n4325K 229\nQ447Q 606\n78TTT 506\n66626 406\n27444 621\n27Q3Q 584\n28T27 964\nQTTTT 598\n8JJA8 649\nQ773K 847\n2AT8Q 231\n39574 241\n34QJ8 239\nJ66KK 98\n54J3A 164\n77A87 582\nKJJKK 768\nJ6892 846\n77TT7 911\n44834 184\nAJKJT 329\nTTTA4 596\n5453A 792\n5666J 493\nTA3A8 333\n7TA7A 585\n7JJ77 645\nK6KK7 843\nKKK9K 555\n33KT6 123\n5AA84 958\n93798 341\n555AA 783\nTAT33 266\nQ6QQ6 940\n7AJQ6 205\nA6656 572\nJ8JJJ 660\n77477 861\nQJ666 173\nQ344J 955\n6699A 332\n529A9 976\nQK6KQ 640\n44QQQ 298\nJ399K 9\n7J294 966\n9TAT7 613\n2K3Q7 680\n999Q5 532\nT899T 198\nQ95QK 42\n97947 509\nK777K 682\nK45K4 554\n82333 153\n2J7KK 802\nQ6999 187\n27732 41\n57865 628\n62226 47\nQQKA2 398\n77772 129\n67499 727\n3J8K8 752\n5J27A 833\n84444 741\n8KJKK 567\nAAJJK 904\n87768 780\n66766 534\nAA66A 989\n76868 32\n47TTT 512\n55J53 748\n49249 550\n2566K 812\n5775A 376\nJ9K99 399\n9J66J 325\n36JQ6 364\n99979 909\n2A792 608\nJ848K 326\n73374 533\n3J655 591\nQ6QQ4 580\n65656 845\n54477 626\n6299Q 225\nQ8768 707\n7989K 408\n7QQ73 525\n76Q2A 653\nA4A44 369\n23737 618\n49696 24\n52542 905\n22262 31\nK686K 346\n5TT8J 185\n76TJ6 494\nTQJ77 711\n33636 120\n4TA97 662\n43946 691\n77262 986\n7J7KJ 370\nTTK9Q 391\nT8387 498\n43J79 781\n66846 133\n2J222 765\n3KTQ3 222\n26KQJ 641\nT94KQ 857\nKKK7K 918\n4644Q 152\n77K6K 282\nA4444 664\nQQ8AA 644\n555JJ 997\n22727 824\nK3K2K 929\n9QK88 696\n44465 480\nK3JK3 110\n22242 718\nKKK9J 998\nQTQTT 462\nTTT6T 214\n5QQTQ 8\nKT593 789\n8QQ68 181\nQKK44 357\nJJQ7J 587\n8858J 522\n55559 351\n93995 673\n9AAT9 700\n78JQJ 739\n32323 950\n3338A 747\nJ22TT 803\nT2893 475\n3T36T 292\nJQ5T2 759\n274Q4 142\n7737Q 579\n92T22 343\nQ2QQ8 100\n762J8 309\nA288A 322\n3569Q 932\n59455 972\n6J668 836\n29A84 562\nQQAQA 468\n8QJKA 76\n66JK5 85\nTTT28 634\n2K67T 146\n44A42 674\nK97KT 40\n55942 839\nJQ5QT 306\n22TTT 274\nQ3J6A 61\n65555 204\n9T9T9 260\nJ6778 269\n7727J 466\n72T58 995\n3333K 910\n49292 108\n9A292 520\nT7TTK 527\n44888 209\nQT6Q3 848\n2JQ22 354\n73TT3 588\n6KT83 507\nAA8AA 519\nQK4QK 558\n9KA97 390\n6K2KK 526\n47KJ2 590\n666QQ 253\n4AQ8J 88\nAJA82 445\n68AAA 652\n6A42K 64\nK2225 237\n6544K 334\n3685A 178\n44324 797\nQ27K6 687\nTTT42 614\n99595 384\n74474 412\nJ5745 293\nQ4444 864"

var LABELS = map[string]int{"2": 1, "3": 2, "4": 3, "5": 4, "6": 5, "7": 6, "8": 7, "9": 8, "T": 9, "J": 10, "Q": 11, "K": 12, "A": 13}

var LABELS_WITH_JOKER = map[string]int{"J": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "T": 10, "Q": 11, "K": 12, "A": 13}

func setPowerWithJoker(hand Hand) Hand {
	handWithPower := setPower(hand)
	jockerCount := countJoker(hand.cards)
	switch handWithPower.power {
	case 7: // 5 dello stesso tipo
	case 6: // poker
		if jockerCount > 0 {
			handWithPower.power = 7
		}
	case 5: // full house
		if jockerCount > 0 {
			handWithPower.power = 7
		}
	case 4: // tripletta
		switch jockerCount {
		case 0:
			handWithPower.power = 4
		case 1: // 333J2 --> poker
			handWithPower.power = 6
		case 2: // 33JJ3 --> 5s
			handWithPower.power = 7
		case 3: // JJJ23 --> poker
			handWithPower.power = 6
		}
	case 3: // doppia coppia
		if jockerCount == 1 {
			handWithPower.power = 5
		} else if jockerCount == 2 {
			handWithPower.power = 6
		}

	case 2: // coppia
		if jockerCount == 1 || jockerCount == 2 {
			handWithPower.power = 4
		} else if jockerCount == 3 {
			handWithPower.power = 5
		}
	case 1: // tutti diversi
		if jockerCount > 0 {
			handWithPower.power += 1

		}
	}

	return handWithPower
}

func countJoker(cards string) int {
	var jokers int
	for _, item := range cards {
		if item == 'J' {
			jokers++
		}
	}
	return jokers
}

func setPower(hand Hand) Hand {
	pattern := countCards(hand.cards)
	// Five of a kind
	if reflect.DeepEqual(pattern, []int{5}) {
		hand.power = 7
		// Four of a kind
	} else if reflect.DeepEqual(pattern, []int{1, 4}) || reflect.DeepEqual(pattern, []int{4, 1}) {
		hand.power = 6
		//  Full house
	} else if reflect.DeepEqual(pattern, []int{3, 2}) || reflect.DeepEqual(pattern, []int{2, 3}) {
		hand.power = 5
		// Three of a kind
	} else if reflect.DeepEqual(pattern, []int{3, 1, 1}) || reflect.DeepEqual(pattern, []int{1, 3, 1}) || reflect.DeepEqual(pattern, []int{1, 1, 3}) {
		hand.power = 4
		//  Two pair
	} else if reflect.DeepEqual(pattern, []int{2, 2, 1}) || reflect.DeepEqual(pattern, []int{1, 2, 2}) || reflect.DeepEqual(pattern, []int{2, 1, 2}) {
		hand.power = 3
		// one pair
	} else if reflect.DeepEqual(pattern, []int{2, 1, 1, 1}) || reflect.DeepEqual(pattern, []int{1, 2, 1, 1}) || reflect.DeepEqual(pattern, []int{1, 1, 2, 1}) || reflect.DeepEqual(pattern, []int{1, 1, 1, 2}) {
		hand.power = 2
	} else {
		hand.power = 1
	}
	return hand
}

func countCards(input string) []int {
	counter := make(map[string]int)
	for _, char := range input {
		counter[string(char)] += 1
	}
	var result []int
	for _, v := range counter {
		result = append(result, v)
	}
	return result
}

func Compute(inputString string) int {
	hands := make([]Hand, 0)
	for _, input := range strings.Split(inputString, "\n") {
		v := strings.Split(input, " ")
		bid, _ := strconv.Atoi(v[1])
		hand := setPower(Hand{bid: bid, cards: v[0]})
		hands = append(hands, hand)
	}
	computeRank(hands, LABELS)
	return computeWinnings(hands)
}

type lessFunc func(p1, p2 *Hand) int

type multiSorter struct {
	hands []Hand
	less  []lessFunc
}

// Sort sorts the argument slice according to the less functions passed to OrderedBy.
func (ms *multiSorter) Sort(hands []Hand) {
	ms.hands = hands
	sort.Sort(ms)
}

// OrderedBy returns a Sorter that sorts using the less functions, in order.
// Call its Sort method to sort the data.
func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

// Len is part of sort.Interface.
func (ms *multiSorter) Len() int {
	return len(ms.hands)
}

func (ms *multiSorter) Swap(i, j int) {
	ms.hands[i], ms.hands[j] = ms.hands[j], ms.hands[i]
}

// Less is part of sort.Interface. It is implemented by looping along the
// less functions until it finds a comparison that discriminates between
// the two items
func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.hands[i], &ms.hands[j]
	// Try all but the last comparison.
	var k int
	for k = 0; k < len(ms.less); k++ {
		less := ms.less[k]
		switch less(p, q) {
		case -1:
			// p < q, so we have a decision.
			return true
		case 1:
			// p > q, so we have a decision.
			return false
		}
		// p == q; try the next comparison.
	}
	// all comparison resulted in equality, so return false (equal is not less than)
	return false
}

func computeRank(hands []Hand, labelsMap map[string]int) {
	power := func(p1, p2 *Hand) int {
		if p1.power < p2.power {
			return -1
		} else if p1.power == p2.power {
			return 0
		} else {
			return 1
		}
	}
	labels := func(p1, p2 *Hand) int {
		for index := range p1.cards {
			if labelsMap[string(p1.cards[index])] < labelsMap[string(p2.cards[index])] {
				return -1
			} else if labelsMap[string(p1.cards[index])] > labelsMap[string(p2.cards[index])] {
				return 1
			}
		}
		return 0
	}
	OrderedBy(power, labels).Sort(hands)
}

func computeWinnings(hands []Hand) int {
	var total int
	for i, hand := range hands {
		total += hand.bid * (i + 1)
	}
	return total
}

// out : 250347426

func Run() {
	hands := make([]Hand, 0)
	for _, input := range strings.Split(INPUT, "\n") {
		v := strings.Split(input, " ")
		bid, _ := strconv.Atoi(v[1])
		hand := setPower(Hand{bid: bid, cards: v[0]})
		hands = append(hands, hand)
	}
	computeRank(hands, LABELS)
	// fmt.Printf("%v\n", hands)
	println(computeWinnings(hands))
}

func RunWithJoker(input string) int {
	if input == "" {
		input = INPUT
	}
	hands := make([]Hand, 0)
	for _, input := range strings.Split(input, "\n") {
		v := strings.Split(input, " ")
		bid, _ := strconv.Atoi(v[1])
		hand := setPowerWithJoker(Hand{bid: bid, cards: v[0]})
		hands = append(hands, hand)
	}
	computeRank(hands, LABELS_WITH_JOKER)
	// fmt.Printf("%v\n", hands)
	return computeWinnings(hands)
}
